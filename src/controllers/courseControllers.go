package controllers

import (
	"log"
	"strconv"

	"segmenta/src/models"
	"segmenta/src/repositories"
	"segmenta/src/utils"

	"github.com/gin-gonic/gin"
)

func GetAllCourses(c *gin.Context) {
	userID, ok := utils.GetUserID(c, "[GET COURSES]")
	if !ok {
		return
	}

	courses, errorHandler := repositories.GetAllCourses(userID)
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[GET COURSES] Error fetching courses", 500)
		return
	}

	utils.SendSuccessResponse(c, "[GET COURSES] Courses fetched successfully", gin.H{"courses": courses})
}

func GetCourseByID(c *gin.Context) {
	userID, ok := utils.GetUserID(c, "[GET COURSE BY ID]")
	if !ok {
		return
	}

	courseIDStr := c.Param("id")
	courseID, errorHandler := strconv.ParseUint(courseIDStr, 10, 64)
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[GET COURSE BY ID] Invalid course ID", 400)
		return
	}

	course, errorHandler := repositories.GetCourseByID(uint(courseID))
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[GET COURSE BY ID] Error fetching course", 500)
		return
	}

	if uint(course.UserID) != userID {
		utils.SendErrorResponse(c, "[GET COURSE BY ID] Forbidden", 403)
		return
	}

	utils.SendSuccessResponse(c, "[GET COURSE BY ID] Course fetched successfully", gin.H{"course": course})
}

func CreateCourse(c *gin.Context) {
	userID, ok := utils.GetUserID(c, "[CREATE COURSE]")
	if !ok {
		return
	}

	var courseCreateRequest struct {
		VideoLink string `json:"video_link" binding:"required"`
	}

	if err := c.ShouldBindJSON(&courseCreateRequest); err != nil {
		utils.SendErrorResponse(c, "[CREATE COURSE] Invalid request data", 400)
		return
	}

	course := models.Course{
		UserID:    int(userID),
		VideoLink: &courseCreateRequest.VideoLink,
	}

	if err := repositories.CreateCourse(&course); err != nil {
		utils.SendErrorResponse(c, "[CREATE COURSE] Error creating course", 500)
		return
	}

	utils.SendSuccessResponse(c, "[CREATE COURSE] Course created successfully", gin.H{"course": course})
}

func AutoCreateCourses(c *gin.Context) {
	userID, ok := utils.GetUserID(c, "[AUTO CREATE COURSE]")
	if !ok {
		return
	}

	var courseCreateRequest struct {
		VideoLink string `json:"video_link" binding:"required"`
	}

	if errorHandler := c.ShouldBindJSON(&courseCreateRequest); errorHandler != nil {
		utils.SendErrorResponse(c, "[AUTO CREATE COURSE] Invalid request data", 400)
		return
	}

	if utils.LinkDuplicateCheck(c, "[AUTO CREATE COURSE]", "course_link", courseCreateRequest.VideoLink) {
		return
	}

	metadata, errorHandler := utils.FetchVideoMetadata(courseCreateRequest.VideoLink)
	if errorHandler != nil {
		log.Printf("[AUTO CREATE COURSE] FetchVideoMetadata error: %v", errorHandler)
		course := models.Course{
			UserID:    int(userID),
			VideoLink: &courseCreateRequest.VideoLink,
		}
		if createErr := repositories.CreateCourse(&course); createErr != nil {
			utils.SendErrorResponse(c, "[AUTO CREATE COURSE] Error creating course", 500)
			return
		}
		utils.SendSuccessResponse(c, "[AUTO CREATE COURSE] Course created (metadata fetch failed, only video link saved)", gin.H{"course": course, "chapters": []models.Chapter{}})
		return
	}

	//Create course with full metadata
	course := models.Course{
		UserID:            int(userID),
		Title:             metadata.Title,
		Description:       &metadata.Description,
		Channel:           &metadata.Channel,
		ChannelLink:       &metadata.ChannelLink,
		VideoLink:         &metadata.VideoLink,
		ThumbnailImageURL: &metadata.ThumbnailImageURL,
	}

	if errorHandler := repositories.CreateCourse(&course); errorHandler != nil {
		utils.SendErrorResponse(c, "[AUTO CREATE COURSE] Error creating course", 500)
		return
	}

	// Auto-create chapters from description timestamps
	var createdChapters []models.Chapter
	if utils.IsDurationInDescription(metadata.Description) {
		chapters := utils.ChapterMakerFromDescription(metadata.Description, int(course.CourseID))
		for _, chapter := range chapters {
			if errorHandler := repositories.CreateChapter(&chapter); errorHandler == nil {
				createdChapters = append(createdChapters, chapter)
			}
		}
	}

	utils.SendSuccessResponse(c, "[AUTO CREATE COURSE] Course and chapters created successfully", gin.H{
		"course":   course,
		"chapters": createdChapters,
	})
}

func UpdateCourse(c *gin.Context) {
	userID, ok := utils.GetUserID(c, "[UPDATE COURSE]")
	if !ok {
		return
	}

	courseIDStr := c.Param("id")
	courseID, err := strconv.ParseUint(courseIDStr, 10, 64)
	if err != nil {
		utils.SendErrorResponse(c, "[UPDATE COURSE] Invalid course ID", 400)
		return
	}

	existingCourse, errorHandler := repositories.GetCourseByID(uint(courseID))
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[UPDATE COURSE] Error fetching course", 500)
		return
	}

	if uint(existingCourse.UserID) != userID {
		utils.SendErrorResponse(c, "[UPDATE COURSE] Forbidden", 403)
		return
	}

	var courseUpdateRequest models.Course
	if err := c.ShouldBindJSON(&courseUpdateRequest); err != nil {
		utils.SendErrorResponse(c, "[UPDATE COURSE] Invalid request data", 400)
		return
	}

	if err := repositories.UpdateCourseByID(uint(courseID), &courseUpdateRequest); err != nil {
		utils.SendErrorResponse(c, "[UPDATE COURSE] Error updating course", 500)
		return
	}

	utils.SendSuccessResponse(c, "[UPDATE COURSE] Course updated successfully", nil)
}

func DeleteCourse(c *gin.Context) {
	userID, ok := utils.GetUserID(c, "[DELETE COURSE]")
	if !ok {
		return
	}

	courseIDStr := c.Param("id")
	courseID, errorHandler := strconv.ParseUint(courseIDStr, 10, 64)
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[DELETE COURSE] Invalid course ID", 400)
		return
	}

	existingCourse, errorHandler := repositories.GetCourseByID(uint(courseID))
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[DELETE COURSE] Error fetching course", 500)
		return
	}

	if uint(existingCourse.UserID) != userID {
		utils.SendErrorResponse(c, "[DELETE COURSE] Forbidden", 403)
		return
	}

	// Check if theres chapters associated with the course, if yes delete them first
	chapters, errorHandler := repositories.GetAllChaptersByCourseID(uint(courseID))
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[DELETE COURSE] Error fetching chapters", 500)
		return
	}
	
	for _, chapter := range chapters {
		if errorHandler := repositories.DeleteChapterByID(chapter.ChapterID); errorHandler != nil {
			utils.SendErrorResponse(c, "[DELETE COURSE] Error deleting chapter", 500)
			return
		}
	}

	if errorHandler := repositories.DeleteCourseByID(uint(courseID)); errorHandler != nil {
		utils.SendErrorResponse(c, "[DELETE COURSE] Error deleting course", 500)
		return
	}

	utils.SendSuccessResponse(c, "[DELETE COURSE] Course deleted successfully", nil)
}

// Chapter Controllers
func GetAllChaptersByCourseID(c *gin.Context) {
	userID, ok := utils.GetUserID(c, "[GET CHAPTERS BY COURSE ID]")
	if !ok {
		return
	}

	courseIDStr := c.Param("course_id")
	courseID, errorHandler := strconv.ParseUint(courseIDStr, 10, 64)
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[GET CHAPTERS BY COURSE ID] Invalid course ID", 400)
		return
	}

	course, errorHandler := repositories.GetCourseByID(uint(courseID))
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[GET CHAPTERS BY COURSE ID] Error fetching course", 500)
		return
	}

	if uint(course.UserID) != userID {
		utils.SendErrorResponse(c, "[GET CHAPTERS BY COURSE ID] Forbidden", 403)
		return
	}

	chapters, errorHandler := repositories.GetAllChaptersByCourseID(uint(courseID))
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[GET CHAPTERS BY COURSE ID] Error fetching chapters", 500)
		return
	}

	utils.SendSuccessResponse(c, "[GET CHAPTERS BY COURSE ID] Chapters fetched successfully", gin.H{"chapters": chapters})
}

func GetOneChapterByID(c *gin.Context) {
	userID, ok := utils.GetUserID(c, "[GET CHAPTER BY ID]")
	if !ok {
		return
	}

	chapterIDStr := c.Param("id")
	chapterID, errorHandler := strconv.ParseUint(chapterIDStr, 10, 64)
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[GET CHAPTER BY ID] Invalid chapter ID", 400)
		return
	}

	chapter, errorHandler := repositories.GetChapterByID(uint(chapterID))
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[GET CHAPTER BY ID] Error fetching chapter", 500)
		return
	}

	// Verify ownership through the parent course
	course, errorHandler := repositories.GetCourseByID(uint(chapter.CourseID))
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[GET CHAPTER BY ID] Error fetching course", 500)
		return
	}

	if uint(course.UserID) != userID {
		utils.SendErrorResponse(c, "[GET CHAPTER BY ID] Forbidden", 403)
		return
	}

	utils.SendSuccessResponse(c, "[GET CHAPTER BY ID] Chapter fetched successfully", gin.H{"chapter": chapter})
}

func CreateChapter(c *gin.Context) {
	userID, ok := utils.GetUserID(c, "[CREATE CHAPTER]")
	if !ok {
		return
	}

	courseIDStr := c.Param("course_id")
	courseID, errorHandler := strconv.ParseUint(courseIDStr, 10, 64)
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[CREATE CHAPTER] Invalid course ID", 400)
		return
	}

	course, errorHandler := repositories.GetCourseByID(uint(courseID))
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[CREATE CHAPTER] Error fetching course", 500)
		return
	}

	if uint(course.UserID) != userID {
		utils.SendErrorResponse(c, "[CREATE CHAPTER] Forbidden", 403)
		return
	}

	var chapterCreateRequest struct {
		Title          string  `json:"title" binding:"required"`
		VideoTimestamp *string `json:"video_timestamp"`
		Position       int     `json:"position" binding:"required"`
	}

	if err := c.ShouldBindJSON(&chapterCreateRequest); err != nil {
		utils.SendErrorResponse(c, "[CREATE CHAPTER] Invalid request data", 400)
		return
	}

	chapter := models.Chapter{
		CourseID:       int(courseID),
		Title:          chapterCreateRequest.Title,
		VideoTimestamp: chapterCreateRequest.VideoTimestamp,
		Position:       chapterCreateRequest.Position,
	}

	if err := repositories.CreateChapter(&chapter); err != nil {
		utils.SendErrorResponse(c, "[CREATE CHAPTER] Error creating chapter", 500)
		return
	}

	utils.SendSuccessResponse(c, "[CREATE CHAPTER] Chapter created successfully", gin.H{"chapter": chapter})
}

func UpdateChapter(c *gin.Context) {
	userID, ok := utils.GetUserID(c, "[UPDATE CHAPTER]")
	if !ok {
		return
	}

	chapterIDStr := c.Param("id")
	chapterID, errorHandler := strconv.ParseUint(chapterIDStr, 10, 64)
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[UPDATE CHAPTER] Invalid chapter ID", 400)
		return
	}

	existingChapter, errorHandler := repositories.GetChapterByID(uint(chapterID))
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[UPDATE CHAPTER] Error fetching chapter", 500)
		return
	}

	course, errorHandler := repositories.GetCourseByID(uint(existingChapter.CourseID))
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[UPDATE CHAPTER] Error fetching course", 500)
		return
	}

	if uint(course.UserID) != userID {
		utils.SendErrorResponse(c, "[UPDATE CHAPTER] Forbidden", 403)
		return
	}

	var chapterUpdateRequest models.Chapter
	if err := c.ShouldBindJSON(&chapterUpdateRequest); err != nil {
		utils.SendErrorResponse(c, "[UPDATE CHAPTER] Invalid request data", 400)
		return
	}

	if err := repositories.UpdateChapterByID(uint(chapterID), &chapterUpdateRequest); err != nil {
		utils.SendErrorResponse(c, "[UPDATE CHAPTER] Error updating chapter", 500)
		return
	}

	utils.SendSuccessResponse(c, "[UPDATE CHAPTER] Chapter updated successfully", nil)
}

func DeleteChapter(c *gin.Context) {
	userID, ok := utils.GetUserID(c, "[DELETE CHAPTER]")
	if !ok {
		return
	}

	chapterIDStr := c.Param("id")
	chapterID, errorHandler := strconv.ParseUint(chapterIDStr, 10, 64)
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[DELETE CHAPTER] Invalid chapter ID", 400)
		return
	}

	existingChapter, errorHandler := repositories.GetChapterByID(uint(chapterID))
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[DELETE CHAPTER] Error fetching chapter", 500)
		return
	}

	course, errorHandler := repositories.GetCourseByID(uint(existingChapter.CourseID))
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[DELETE CHAPTER] Error fetching course", 500)
		return
	}

	if uint(course.UserID) != userID {
		utils.SendErrorResponse(c, "[DELETE CHAPTER] Forbidden", 403)
		return
	}

	if err := repositories.DeleteChapterByID(uint(chapterID)); err != nil {
		utils.SendErrorResponse(c, "[DELETE CHAPTER] Error deleting chapter", 500)
		return
	}

	utils.SendSuccessResponse(c, "[DELETE CHAPTER] Chapter deleted successfully", nil)
}

// func MakePublicCourse(c *gin.Context) {
// 	userID, ok := utils.GetUserID(c, "[MAKE PUBLIC COURSE]")
// 	if !ok {
// 		return
// 	}
	
// 	courseIDStr := c.Param("id")
// 	courseID, errorHandler := strconv.ParseUint(courseIDStr, 10, 64)
// 	if errorHandler != nil {
// 		utils.SendErrorResponse(c, "[MAKE PUBLIC COURSE] Invalid course ID", 400)
// 		return
// 	}

// 	existingCourse, errorHandler := repositories.GetCourseByID(uint(courseID))
// 	if errorHandler != nil {
// 		utils.SendErrorResponse(c, "[MAKE PUBLIC COURSE] Error fetching course", 500)
// 		return
// 	}
	
// 	if uint(existingCourse.UserID) != userID {
// 		utils.SendErrorResponse(c, "[MAKE PUBLIC COURSE] Forbidden", 403)
// 		return
// 	}

// 	existingCourse.IsPublic = true
// }