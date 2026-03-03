package controllers

import (
	"strconv"

	"segmenta/src/models"
	"segmenta/src/repositories"
	"segmenta/src/utils"

	"github.com/gin-gonic/gin"
)

func GetAllCourses(c *gin.Context) {
	userID, ok := utils.getUserID(c, "[GET COURSES]")
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
	userID, ok := utils.getUserID(c, "[GET COURSE BY ID]")
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
	userID, ok := utils.getUserID(c, "[CREATE COURSE]")
	if !ok {
		return
	}

	var courseCreateRequest struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description" binding:"required"`
		Link        string `json:"link" binding:"required"`
	}

	if err := c.ShouldBindJSON(&courseCreateRequest); err != nil {
		utils.SendErrorResponse(c, "[CREATE COURSE] Invalid request data", 400)
		return
	}

	desc := courseCreateRequest.Description
	course := models.Course{
		UserID:      int(userID),
		Title:       courseCreateRequest.Title,
		Description: &desc,
		Link:        &courseCreateRequest.Link,
	}

	if err := repositories.CreateCourse(&course); err != nil {
		utils.SendErrorResponse(c, "[CREATE COURSE] Error creating course", 500)
		return
	}

	utils.SendSuccessResponse(c, "[CREATE COURSE] Course created successfully", gin.H{"course": course})
}

func UpdateCourse(c *gin.Context) {
	userID, ok := utils.getUserID(c, "[UPDATE COURSE]")
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
	userID, ok := utils.getUserID(c, "[DELETE COURSE]")
	if !ok {
		return
	}

	courseIDStr := c.Param("id")
	courseID, err := strconv.ParseUint(courseIDStr, 10, 64)
	if err != nil {
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

	if err := repositories.DeleteCourseByID(uint(courseID)); err != nil {
		utils.SendErrorResponse(c, "[DELETE COURSE] Error deleting course", 500)
		return
	}

	utils.SendSuccessResponse(c, "[DELETE COURSE] Course deleted successfully", nil)
}