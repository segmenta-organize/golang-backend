package services

import (
	"log"
	"strconv"

	"segmenta/src/models"
	"segmenta/src/repositories"
	"segmenta/src/utils"

	"github.com/gin-gonic/gin"
)

func GetAllEnrolledCourses(c *gin.Context) {
	userID, ok := utils.GetUserID(c, "[GET ALL ENROLLED COURSES]")
	if !ok {
		return
	}

	courses, errorHandler := repositories.GetAllEnrolledCourses(userID)
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[GET ALL ENROLLED COURSES] Error fetching courses", 500)
		return
	}

	utils.SendSuccessResponse(c, "[GET ALL ENROLLED COURSES] Courses fetched successfully", gin.H{"courses": courses})
}

func GetOneCourseWithChaptersByID(c *gin.Context) {
	userID, ok := utils.GetUserID(c, "[GET ONE COURSE WITH CHAPTERS BY ID]")
	if !ok {
		return
	}

	courseIDStr := c.Param("id")
	courseID, errorHandler := strconv.ParseUint(courseIDStr, 10, 64)
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[GET ONE COURSE WITH CHAPTERS BY ID] Invalid course ID", 400)
		return
	}

	course, errorHandler := repositories.GetOneCourseByID(uint(courseID))
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[GET ONE COURSE WITH CHAPTERS BY ID] Error fetching course", 500)
		return
	}

	if uint(course.UserID) != userID {
		utils.SendErrorResponse(c, "[GET ONE COURSE WITH CHAPTERS BY ID] Forbidden", 403)
		return
	}

	currentChapters, errorHandler := repositories.GetAllChaptersByCourseID(uint(courseID))
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[GET ONE COURSE WITH CHAPTERS BY ID] Error fetching chapter", 500)
		return
	}

	utils.SendSuccessResponse(c, "[GET ONE COURSE BY ID] Course fetched successfully", gin.H{
		"course": course, 
		"chapters": currentChapters,
	})
}

func CreateManualCourseWithChapters(c *gin.Context) {
	userID, ok := utils.GetUserID(c, "[CREATE COURSE]")
	if !ok {
		return
	}

	type chapterCreateRequest struct {
		Title          string  `json:"title" binding:"required"`
		VideoTimestamp *string `json:"video_timestamp,omitempty"`
	}

	var request struct {
		Title             string                 `json:"title" binding:"required"`
		Description       *string                `json:"description,omitempty"`
		Channel           *string                `json:"channel,omitempty"`
		ChannelLink       *string                `json:"channel_link,omitempty"`
		VideoLink         *string                `json:"video_link" binding:"required"`
		ThumbnailImageURL *string                `json:"thumbnail_image_url,omitempty"`
		Chapters          []chapterCreateRequest `json:"chapters,omitempty"`
	}

	if errorHandler := c.ShouldBindJSON(&request); errorHandler != nil {
		utils.SendErrorResponse(c, "[CREATE COURSE] Invalid request data", 400)
		return
	}

	if request.VideoLink != nil && *request.VideoLink != "" {
		if utils.LinkDuplicateCheck(c, "[CREATE COURSE]", "course_link", *request.VideoLink) {
			return
		}
	}

	course := models.Course{
		UserID:            int(userID),
		Title:             request.Title,
		Description:       request.Description,
		Channel:           request.Channel,
		ChannelLink:       request.ChannelLink,
		VideoLink:         request.VideoLink,
		ThumbnailImageURL: request.ThumbnailImageURL,
	}

	if errorHandler := repositories.CreateCourse(&course); errorHandler != nil {
		utils.SendErrorResponse(c, "[CREATE COURSE] Error creating course", 500)
		return
	}

	createdChapters := make([]models.Chapter, 0, len(request.Chapters))
	for i, chapterRequest := range request.Chapters {
		chapter := models.Chapter{
			CourseID:       int(course.CourseID),
			Title:          chapterRequest.Title,
			VideoTimestamp: chapterRequest.VideoTimestamp,
			Position:       i + 1,
		}

		if errorHandler := repositories.CreateChapter(&chapter); errorHandler != nil {
			_ = repositories.DeleteChaptersByCourseID(uint(course.CourseID))
			_ = repositories.DeleteOneCourseByID(uint(course.CourseID))

			utils.SendErrorResponse(c, "[CREATE COURSE] Error creating chapters", 500)
			return
		}

		createdChapters = append(createdChapters, chapter)
	}

	utils.SendSuccessResponse(c, "[CREATE COURSE] Course created successfully", gin.H{
		"course":   course,
		"chapters": createdChapters,
	})
}

func AutoCreateCourseWithChapters(c *gin.Context) {
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

func UpdateCourseWithChapters(c *gin.Context) {
	userID, ok := utils.GetUserID(c, "[UPDATE COURSE]")
	if !ok {
		return
	}

	courseIDStr := c.Param("id")
	courseID, errorHandler := strconv.ParseUint(courseIDStr, 10, 64)
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[UPDATE COURSE] Invalid course ID", 400)
		return
	}

	existingCourse, errorHandler := repositories.GetOneCourseByID(uint(courseID))
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[UPDATE COURSE] Error fetching course", 500)
		return
	}

	if uint(existingCourse.UserID) != userID {
		utils.SendErrorResponse(c, "[UPDATE COURSE] Forbidden", 403)
		return
	}

	type chapterUpsertRequest struct {
		ChapterID       *uint   `json:"chapter_id,omitempty"`
		Title           string  `json:"title,omitempty"`
		VideoTimestamp  *string `json:"video_timestamp,omitempty"`
		Position        int     `json:"position,omitempty"`
		Delete          bool    `json:"delete,omitempty"`
	}

	type courseUpdateWithChaptersRequest struct {
		Title             *string                `json:"title,omitempty"`
		Description       *string                `json:"description,omitempty"`
		Channel           *string                `json:"channel,omitempty"`
		ChannelLink       *string                `json:"channel_link,omitempty"`
		VideoLink         *string                `json:"video_link,omitempty"`
		ThumbnailImageURL *string                `json:"thumbnail_image_url,omitempty"`
		Chapters          []chapterUpsertRequest `json:"chapters,omitempty"`
	}

	var request courseUpdateWithChaptersRequest
	if errorHandler = c.ShouldBindJSON(&request); errorHandler != nil {
		utils.SendErrorResponse(c, "[UPDATE COURSE] Invalid request data", 400)
		return
	}

	courseUpdate := models.Course{}
	if request.Title != nil {
		courseUpdate.Title = *request.Title
	}
	if request.Description != nil {
		courseUpdate.Description = request.Description
	}
	if request.Channel != nil {
		courseUpdate.Channel = request.Channel
	}
	if request.ChannelLink != nil {
		courseUpdate.ChannelLink = request.ChannelLink
	}
	if request.VideoLink != nil {
		if *request.VideoLink != "" && existingCourse.VideoLink != nil && *existingCourse.VideoLink != *request.VideoLink {
			if utils.LinkDuplicateCheck(c, "[UPDATE COURSE]", "course_link", *request.VideoLink) {
				return
			}
		}
		courseUpdate.VideoLink = request.VideoLink
	}
	if request.ThumbnailImageURL != nil {
		courseUpdate.ThumbnailImageURL = request.ThumbnailImageURL
	}

	if request.Chapters != nil {
		if errorHandler = repositories.UpdateCourseByID(uint(courseID), &courseUpdate); errorHandler != nil {
			utils.SendErrorResponse(c, "[UPDATE COURSE] Error updating course", 500)
			return
		}
	
		if errorHandler = repositories.DeleteAllChaptersInOneCourseByID(uint(courseID)); errorHandler != nil {
			utils.SendErrorResponse(c, "[UPDATE COURSE] Error deleting chapters", 500)
			return
		}

		for _, chapterRequest := range request.Chapters {
			if chapterRequest.Delete {
				continue
			}

			chapter := models.Chapter{
				CourseID:       int(courseID),
				Title:          chapterRequest.Title,
				VideoTimestamp: chapterRequest.VideoTimestamp,
				Position:       chapterRequest.Position,
			}

			if errorHandler = repositories.CreateChapter(&chapter); errorHandler != nil {
				utils.SendErrorResponse(c, "[UPDATE COURSE] Error creating chapter", 500)
				return
			}
		}
	}

	updatedCourse, errorHandler := repositories.GetOneCourseByID(uint(courseID))
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[UPDATE COURSE] Error fetching updated course", 500)
		return
	}

	updatedChapters, errorHandler := repositories.GetAllChaptersByCourseID(uint(courseID))
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[UPDATE COURSE] Error fetching updated chapters", 500)
		return
	}

	utils.SendSuccessResponse(c, "[UPDATE COURSE] Course updated successfully", gin.H{
		"course":   updatedCourse,
		"chapters": updatedChapters,
	})
}

func AutoUpdateCourseWithChapters(c *gin.Context) {
	userID, ok := utils.GetUserID(c, "[AUTO UPDATE COURSE]")
	if !ok {
		return
	}

	courseIDStr := c.Param("id")
	courseID, errorHandler := strconv.ParseUint(courseIDStr, 10, 64)
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[AUTO UPDATE COURSE] Invalid course ID", 400)
		return
	}

	existingCourse, errorHandler := repositories.GetOneCourseByID(uint(courseID))
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[AUTO UPDATE COURSE] Error fetching course", 500)
		return
	}

	if uint(existingCourse.UserID) != userID {
		utils.SendErrorResponse(c, "[AUTO UPDATE COURSE] Forbidden", 403)
		return
	}

	var request struct {
		Text 	  string `json:"text" binding:"required"`
		VideoLink *string `json:"video_link,omitempty" binding:"omitempty,url,required_with=Text"`
	}

	if errorHandler = c.ShouldBindJSON(&request); errorHandler != nil {
		utils.SendErrorResponse(c, "[AUTO UPDATE COURSE] Invalid request data", 400)
		return
	}

	if request.VideoLink != nil {
		if *request.VideoLink != "" && existingCourse.VideoLink != nil && *existingCourse.VideoLink != *request.VideoLink {
			if utils.LinkDuplicateCheck(c, "[AUTO UPDATE COURSE]", "course_link", *request.VideoLink) {
				return
			}
		}
	}

	courseUpdate := models.Course{}
	if request.VideoLink != nil {
		metadata := utils.AutoUpdateMetadata(&courseUpdate, *request.VideoLink)
		if metadata != nil {
			utils.SendErrorResponse(c, "[AUTO UPDATE COURSE] Error updating course", 500)
			return
		}
	}

	if errorHandler = repositories.UpdateCourseByID(uint(courseID), &courseUpdate); errorHandler != nil {
		utils.SendErrorResponse(c, "[AUTO UPDATE COURSE] Error updating course", 500)
		return
	}

	if errorHandler = repositories.DeleteAllChaptersInOneCourseByID(uint(courseID)); errorHandler != nil {
		utils.SendErrorResponse(c, "[AUTO UPDATE COURSE] Error deleting chapters", 500)
		return
	}

	chapters := utils.ChapterMakerFromDescription(request.Text, int(courseID))
	var createdChapters []models.Chapter
	for _, chapter := range chapters {
		if errorHandler = repositories.CreateChapter(&chapter); errorHandler == nil {
			createdChapters = append(createdChapters, chapter)
		}
	}

	updatedCourse, errorHandler := repositories.GetOneCourseByID(uint(courseID))
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[AUTO UPDATE COURSE] Error fetching updated course", 500)
		return
	}

	updatedChapters, errorHandler := repositories.GetAllChaptersByCourseID(uint(courseID))
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[AUTO UPDATE COURSE] Error fetching updated chapters", 500)
		return
	}

	utils.SendSuccessResponse(c, "[AUTO UPDATE COURSE] Course updated successfully", gin.H{
		"course":   updatedCourse,
		"chapters": updatedChapters,
	})
}

func DeleteOneCourseByID(c *gin.Context) {
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

	existingCourse, errorHandler := repositories.GetOneCourseByID(uint(courseID))
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[DELETE COURSE] Error fetching course", 500)
		return
	}

	if uint(existingCourse.UserID) != userID {
		utils.SendErrorResponse(c, "[DELETE COURSE] Forbidden", 403)
		return
	}

	if errorHandler = repositories.DeleteAllChaptersInOneCourseByID(uint(courseID)); errorHandler != nil {
		utils.SendErrorResponse(c, "[DELETE COURSE] Error deleting chapters", 500)
		return
	}

	if errorHandler := repositories.DeleteOneCourseByID(uint(courseID)); errorHandler != nil {
		utils.SendErrorResponse(c, "[DELETE COURSE] Error deleting course", 500)
		return
	}

	utils.SendSuccessResponse(c, "[DELETE COURSE] Course deleted successfully", nil)
}

// Publish Course to Explore

func CreatePublicCourseFromCourse(c *gin.Context) {
	userID, ok := utils.GetUserID(c, "[CREATE PUBLIC COURSE]")
	if !ok {
		return
	}

	courseIDStr := c.Param("id")
	courseID, errorHandler := strconv.ParseUint(courseIDStr, 10, 64)
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[CREATE PUBLIC COURSE] Invalid course ID", 400)
		return
	}

	existingCourse, errorHandler := repositories.GetOneCourseByID(uint(courseID))
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[CREATE PUBLIC COURSE] Error fetching course", 500)
		return
	}

	if uint(existingCourse.UserID) != userID {
		utils.SendErrorResponse(c, "[CREATE PUBLIC COURSE] Forbidden", 403)
		return
	}

	if errorHandler := repositories.CreatePublicCourseFromCourse(uint(courseID), userID); errorHandler != nil {
		utils.SendErrorResponse(c, "[CREATE PUBLIC COURSE] Error creating public course", 500)
		return
	}

	utils.SendSuccessResponse(c, "[CREATE PUBLIC COURSE] Public course created successfully", nil)
}

func UpdatePublicCourseFromCourse(c *gin.Context) {
	userID, ok := utils.GetUserID(c, "[UPDATE PUBLIC COURSE]")
	if !ok {
		return
	}

	courseIDStr := c.Param("id")
	courseID, errorHandler := strconv.ParseUint(courseIDStr, 10, 64)
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[UPDATE PUBLIC COURSE] Invalid course ID", 400)
		return
	}

	existingCourse, errorHandler := repositories.GetOneCourseByID(uint(courseID))
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[UPDATE PUBLIC COURSE] Error fetching course", 500)
		return
	}

	if uint(existingCourse.UserID) != userID {
		utils.SendErrorResponse(c, "[UPDATE PUBLIC COURSE] Forbidden", 403)
		return
	}

	// Need explore_course_id to know which public course to update
	var requestBody struct {
		ExploreCourseID uint `json:"explore_course_id" binding:"required"`
	}
	if errorHandler := c.ShouldBindJSON(&requestBody); errorHandler != nil {
		utils.SendErrorResponse(c, "[UPDATE PUBLIC COURSE] explore_course_id is required", 400)
		return
	}

	if errorHandler := repositories.UpdatePublicCourseFromCourse(uint(courseID), requestBody.ExploreCourseID); errorHandler != nil {
		utils.SendErrorResponse(c, "[UPDATE PUBLIC COURSE] Error updating public course", 500)
		return
	}

	utils.SendSuccessResponse(c, "[UPDATE PUBLIC COURSE] Public course updated successfully", nil)
}

// Chapter services

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

	course, errorHandler := repositories.GetOneCourseByID(uint(courseID))
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
	course, errorHandler := repositories.GetOneCourseByID(uint(chapter.CourseID))
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

	course, errorHandler := repositories.GetOneCourseByID(uint(courseID))
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

	if errorHandler := c.ShouldBindJSON(&chapterCreateRequest); errorHandler != nil {
		utils.SendErrorResponse(c, "[CREATE CHAPTER] Invalid request data", 400)
		return
	}

	chapter := models.Chapter{
		CourseID:       int(courseID),
		Title:          chapterCreateRequest.Title,
		VideoTimestamp: chapterCreateRequest.VideoTimestamp,
		Position:       chapterCreateRequest.Position,
	}

	if errorHandler := repositories.CreateChapter(&chapter); errorHandler != nil {
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

	course, errorHandler := repositories.GetOneCourseByID(uint(existingChapter.CourseID))
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[UPDATE CHAPTER] Error fetching course", 500)
		return
	}

	if uint(course.UserID) != userID {
		utils.SendErrorResponse(c, "[UPDATE CHAPTER] Forbidden", 403)
		return
	}

	var chapterUpdateRequest models.Chapter
	if errorHandler := c.ShouldBindJSON(&chapterUpdateRequest); errorHandler != nil {
		utils.SendErrorResponse(c, "[UPDATE CHAPTER] Invalid request data", 400)
		return
	}

	if errorHandler := repositories.UpdateChapterByID(uint(chapterID), &chapterUpdateRequest); errorHandler != nil {
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

	course, errorHandler := repositories.GetOneCourseByID(uint(existingChapter.CourseID))
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[DELETE CHAPTER] Error fetching course", 500)
		return
	}

	if uint(course.UserID) != userID {
		utils.SendErrorResponse(c, "[DELETE CHAPTER] Forbidden", 403)
		return
	}

	if errorHandler := repositories.DeleteChapterByID(uint(chapterID)); errorHandler != nil {
		utils.SendErrorResponse(c, "[DELETE CHAPTER] Error deleting chapter", 500)
		return
	}

	utils.SendSuccessResponse(c, "[DELETE CHAPTER] Chapter deleted successfully", nil)
}

// Chapter → Explore Chapter public functions

func CreatePublicChapterFromChapter(c *gin.Context) {
	userID, ok := utils.GetUserID(c, "[CREATE PUBLIC CHAPTER]")
	if !ok {
		return
	}

	chapterIDStr := c.Param("id")
	chapterID, errorHandler := strconv.ParseUint(chapterIDStr, 10, 64)
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[CREATE PUBLIC CHAPTER] Invalid chapter ID", 400)
		return
	}

	// Verify ownership through course
	chapter, errorHandler := repositories.GetChapterByID(uint(chapterID))
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[CREATE PUBLIC CHAPTER] Chapter not found", 404)
		return
	}

	course, errorHandler := repositories.GetOneCourseByID(uint(chapter.CourseID))
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[CREATE PUBLIC CHAPTER] Course not found", 404)
		return
	}

	if uint(course.UserID) != userID {
		utils.SendErrorResponse(c, "[CREATE PUBLIC CHAPTER] Forbidden", 403)
		return
	}

	var requestBody struct {
		ExploreCourseID uint `json:"explore_course_id" binding:"required"`
	}
	if errorHandler := c.ShouldBindJSON(&requestBody); errorHandler != nil {
		utils.SendErrorResponse(c, "[CREATE PUBLIC CHAPTER] explore_course_id is required", 400)
		return
	}

	if errorHandler := repositories.CreatePublicChapterFromChapter(uint(chapterID), requestBody.ExploreCourseID); errorHandler != nil {
		utils.SendErrorResponse(c, "[CREATE PUBLIC CHAPTER] Error creating public chapter", 500)
		return
	}

	utils.SendSuccessResponse(c, "[CREATE PUBLIC CHAPTER] Public chapter created successfully", nil)
}

func UpdatePublicChapterFromChapter(c *gin.Context) {
	userID, ok := utils.GetUserID(c, "[UPDATE PUBLIC CHAPTER]")
	if !ok {
		return
	}

	chapterIDStr := c.Param("id")
	chapterID, errorHandler := strconv.ParseUint(chapterIDStr, 10, 64)
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[UPDATE PUBLIC CHAPTER] Invalid chapter ID", 400)
		return
	}

	// Verify ownership through course
	chapter, errorHandler := repositories.GetChapterByID(uint(chapterID))
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[UPDATE PUBLIC CHAPTER] Chapter not found", 404)
		return
	}

	course, errorHandler := repositories.GetOneCourseByID(uint(chapter.CourseID))
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[UPDATE PUBLIC CHAPTER] Course not found", 404)
		return
	}

	if uint(course.UserID) != userID {
		utils.SendErrorResponse(c, "[UPDATE PUBLIC CHAPTER] Forbidden", 403)
		return
	}

	var requestBody struct {
		ExploreChapterID uint `json:"explore_chapter_id" binding:"required"`
	}
	if errorHandler := c.ShouldBindJSON(&requestBody); errorHandler != nil {
		utils.SendErrorResponse(c, "[UPDATE PUBLIC CHAPTER] explore_chapter_id is required", 400)
		return
	}

	if errorHandler := repositories.UpdatePublicChapterFromChapter(uint(chapterID), requestBody.ExploreChapterID); errorHandler != nil {
		utils.SendErrorResponse(c, "[UPDATE PUBLIC CHAPTER] Error updating public chapter", 500)
		return
	}

	utils.SendSuccessResponse(c, "[UPDATE PUBLIC CHAPTER] Public chapter updated successfully", nil)
}

func DeletePublicChapterFromChapter(c *gin.Context) {
	userID, ok := utils.GetUserID(c, "[DELETE PUBLIC CHAPTER]")
	if !ok {
		return
	}

	chapterIDStr := c.Param("id")
	chapterID, errorHandler := strconv.ParseUint(chapterIDStr, 10, 64)
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[DELETE PUBLIC CHAPTER] Invalid chapter ID", 400)
		return
	}

	// Verify ownership through course
	chapter, errorHandler := repositories.GetChapterByID(uint(chapterID))
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[DELETE PUBLIC CHAPTER] Chapter not found", 404)
		return
	}

	course, errorHandler := repositories.GetOneCourseByID(uint(chapter.CourseID))
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[DELETE PUBLIC CHAPTER] Course not found", 404)
		return
	}

	if uint(course.UserID) != userID {
		utils.SendErrorResponse(c, "[DELETE PUBLIC CHAPTER] Forbidden", 403)
		return
	}

	var requestBody struct {
		ExploreChapterID uint `json:"explore_chapter_id" binding:"required"`
	}
	if errorHandler := c.ShouldBindJSON(&requestBody); errorHandler != nil {
		utils.SendErrorResponse(c, "[DELETE PUBLIC CHAPTER] explore_chapter_id is required", 400)
		return
	}

	if errorHandler := repositories.DeletePublicChapterFromChapter(requestBody.ExploreChapterID); errorHandler != nil {
		utils.SendErrorResponse(c, "[DELETE PUBLIC CHAPTER] Error deleting public chapter", 500)
		return
	}

	utils.SendSuccessResponse(c, "[DELETE PUBLIC CHAPTER] Public chapter deleted successfully", nil)
}
