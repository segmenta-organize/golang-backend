package controllers

import (
	"strconv"

	"segmenta/src/models"
	"segmenta/src/repositories"
	"segmenta/src/utils"

	"github.com/gin-gonic/gin"
)

// Explore Course Controllers

func GetAllCoursesForExplore(c *gin.Context) {
	userID, ok := utils.GetUserID(c, "[GET ALL COURSES FOR EXPLORE]")
	if !ok {
		return
	}

	availableCourses, errorHandler := repositories.GetAllAvailableCourses(userID)
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[GET ALL COURSES FOR EXPLORE] Error fetching courses", 500)
		return
	}

	utils.SendSuccessResponse(c, "[GET ALL COURSES FOR EXPLORE] Courses fetched successfully", gin.H{"courses": availableCourses})
}

func GetExploredCourseByID(c *gin.Context) {
	_, ok := utils.GetUserID(c, "[GET EXPLORED COURSE BY ID]")
	if !ok {
		return
	}

	exploredCourseIDStr := c.Param("id")
	exploredCourseID, errorHandler := strconv.ParseUint(exploredCourseIDStr, 10, 64)
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[GET EXPLORED COURSE BY ID] Invalid course ID", 400)
		return
	}

	exploredCourse, errorHandler := repositories.GetExploredCourseByID(uint(exploredCourseID))
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[GET EXPLORED COURSE BY ID] Error fetching course", 500)
		return
	}

	utils.SendSuccessResponse(c, "[GET EXPLORED COURSE BY ID] Course fetched successfully", gin.H{"course": exploredCourse})
}

func SearchCourses(c *gin.Context) {
	_, ok := utils.GetUserID(c, "[SEARCH COURSES]")
	if !ok {
		return
	}

	query := c.Query("q")
	if query == "" {
		utils.SendErrorResponse(c, "[SEARCH COURSES] Query parameter 'q' is required", 400)
		return
	}

	searchResults, errorHandler := repositories.SearchCourses(query)
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[SEARCH COURSES] Error searching courses", 500)
		return
	}

	utils.SendSuccessResponse(c, "[SEARCH COURSES] Courses fetched successfully", gin.H{"courses": searchResults})
}

func EnrollInCourse(c *gin.Context) {
	userID, ok := utils.GetUserID(c, "[ENROLL IN COURSE]")
	if !ok {
		return
	}

	courseIDStr := c.Param("id")
	courseID, errorHandler := strconv.ParseUint(courseIDStr, 10, 64)
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[ENROLL IN COURSE] Invalid course ID", 400)
		return
	}

	errorHandler = repositories.EnrollUserInCourse(userID, uint(courseID))
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[ENROLL IN COURSE] Error enrolling in course", 500)
		return
	}

	utils.SendSuccessResponse(c, "[ENROLL IN COURSE] Enrolled in course successfully", nil)
}

// Explore Course Chapter Controllers

func GetAllExploreChapterByCourseID(c *gin.Context) {
	_, ok := utils.GetUserID(c, "[GET ALL EXPLORE CHAPTERS BY COURSE ID]")
	if !ok {
		return
	}

	exploreCourseIDStr := c.Param("course_id")
	exploreCourseID, errorHandler := strconv.ParseUint(exploreCourseIDStr, 10, 64)
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[GET ALL EXPLORE CHAPTERS BY COURSE ID] Invalid course ID", 400)
		return
	}

	chapters, errorHandler := repositories.GetAllExploreChaptersByCourseID(uint(exploreCourseID))
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[GET ALL EXPLORE CHAPTERS BY COURSE ID] Error fetching chapters", 500)
		return
	}

	utils.SendSuccessResponse(c, "[GET ALL EXPLORE CHAPTERS BY COURSE ID] Chapters fetched successfully", gin.H{"chapters": chapters})
}

func GetOneExploreChapterByID(c *gin.Context) {
	_, ok := utils.GetUserID(c, "[GET ONE EXPLORE CHAPTER BY ID]")
	if !ok {
		return
	}

	exploreChapterIDStr := c.Param("id")
	exploreChapterID, errorHandler := strconv.ParseUint(exploreChapterIDStr, 10, 64)
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[GET ONE EXPLORE CHAPTER BY ID] Invalid chapter ID", 400)
		return
	}

	chapter, errorHandler := repositories.GetOneExploreChapterByID(uint(exploreChapterID))
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[GET ONE EXPLORE CHAPTER BY ID] Error fetching chapter", 500)
		return
	}

	utils.SendSuccessResponse(c, "[GET ONE EXPLORE CHAPTER BY ID] Chapter fetched successfully", gin.H{"chapter": chapter})
}

func CreateExploreChapter(c *gin.Context) {
	_, ok := utils.GetUserID(c, "[CREATE EXPLORE CHAPTER]")
	if !ok {
		return
	}

	exploreCourseIDStr := c.Param("course_id")
	exploreCourseID, errorHandler := strconv.ParseUint(exploreCourseIDStr, 10, 64)
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[CREATE EXPLORE CHAPTER] Invalid course ID", 400)
		return
	}

	// Verify the course exists
	_, errorHandler = repositories.GetExploredCourseByID(uint(exploreCourseID))
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[CREATE EXPLORE CHAPTER] Explore course not found", 404)
		return
	}

	var chapterCreateRequest struct {
		Title       string  `json:"title" binding:"required"`
		Description *string `json:"description"`
		Order       int     `json:"order" binding:"required"`
	}

	if err := c.ShouldBindJSON(&chapterCreateRequest); err != nil {
		utils.SendErrorResponse(c, "[CREATE EXPLORE CHAPTER] Invalid request data", 400)
		return
	}

	chapter := models.ExploreChapter{
		ExploreCourseID: uint(exploreCourseID),
		Title:           chapterCreateRequest.Title,
		Description:     chapterCreateRequest.Description,
		Order:           chapterCreateRequest.Order,
	}

	if err := repositories.CreateExploreChapter(&chapter); err != nil {
		utils.SendErrorResponse(c, "[CREATE EXPLORE CHAPTER] Error creating chapter", 500)
		return
	}

	utils.SendSuccessResponse(c, "[CREATE EXPLORE CHAPTER] Chapter created successfully", gin.H{"chapter": chapter})
}

func UpdateExploreChapter(c *gin.Context) {
	_, ok := utils.GetUserID(c, "[UPDATE EXPLORE CHAPTER]")
	if !ok {
		return
	}

	exploreChapterIDStr := c.Param("id")
	exploreChapterID, errorHandler := strconv.ParseUint(exploreChapterIDStr, 10, 64)
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[UPDATE EXPLORE CHAPTER] Invalid chapter ID", 400)
		return
	}

	_, errorHandler = repositories.GetOneExploreChapterByID(uint(exploreChapterID))
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[UPDATE EXPLORE CHAPTER] Chapter not found", 404)
		return
	}

	var chapterUpdateRequest models.ExploreChapter
	if err := c.ShouldBindJSON(&chapterUpdateRequest); err != nil {
		utils.SendErrorResponse(c, "[UPDATE EXPLORE CHAPTER] Invalid request data", 400)
		return
	}

	if err := repositories.UpdateExploreChapterByID(uint(exploreChapterID), &chapterUpdateRequest); err != nil {
		utils.SendErrorResponse(c, "[UPDATE EXPLORE CHAPTER] Error updating chapter", 500)
		return
	}

	utils.SendSuccessResponse(c, "[UPDATE EXPLORE CHAPTER] Chapter updated successfully", nil)
}

func DeleteExploreChapter(c *gin.Context) {
	_, ok := utils.GetUserID(c, "[DELETE EXPLORE CHAPTER]")
	if !ok {
		return
	}

	exploreChapterIDStr := c.Param("id")
	exploreChapterID, errorHandler := strconv.ParseUint(exploreChapterIDStr, 10, 64)
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[DELETE EXPLORE CHAPTER] Invalid chapter ID", 400)
		return
	}

	_, errorHandler = repositories.GetOneExploreChapterByID(uint(exploreChapterID))
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[DELETE EXPLORE CHAPTER] Chapter not found", 404)
		return
	}

	if err := repositories.DeleteExploreChapterByID(uint(exploreChapterID)); err != nil {
		utils.SendErrorResponse(c, "[DELETE EXPLORE CHAPTER] Error deleting chapter", 500)
		return
	}

	utils.SendSuccessResponse(c, "[DELETE EXPLORE CHAPTER] Chapter deleted successfully", nil)
}
