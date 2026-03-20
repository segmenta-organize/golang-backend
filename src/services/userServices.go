package services

import (
	"strconv"

	"segmenta/src/repositories"
	"segmenta/src/utils"

	"github.com/gin-gonic/gin"
)

func GetOneUserByUserID(c *gin.Context) {
	userIDStr := c.Param("id")

	if userIDStr == "" {
		utils.SendErrorResponse(c, "[GET ONE PROFILE] User ID is required", 400)
		return
	}

	userID, errorHandler := strconv.Atoi(userIDStr)
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[GET ONE PROFILE] Invalid user ID", 400)
		return
	}

	user, errorHandler := repositories.GetUserByUserID(uint(userID))
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[GET ONE PROFILE] User not found", 404)
		return
	}

	utils.SendSuccessResponse(c, "[GET ONE PROFILE] User profile retrieved successfully", user)
}

func UpdateUserByUserID(c *gin.Context) {
	userIDStr := c.Param("id")

	if userIDStr == "" {
		utils.SendErrorResponse(c, "[UPDATE PROFILE] User ID is required", 400)
		return
	}

	userID, errorHandler := strconv.Atoi(userIDStr)
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[UPDATE PROFILE] Invalid user ID", 400)
		return
	}

	user, errorHandler := repositories.GetUserByUserID(uint(userID))
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[UPDATE PROFILE] User not found", 404)
		return
	}

	var updateRequest struct {
		FullName       string `json:"full_name"`
		Email          string `json:"email"`
		Password       string `json:"password"`
		Bio            string `json:"bio"`
		ProfilePicLink string `json:"profile_pic_link"`
	}

	if errorHandler := c.ShouldBindJSON(&updateRequest); errorHandler != nil {
		utils.SendErrorResponse(c, "[UPDATE PROFILE] Invalid request data", 400)
		return
	}

	if updateRequest.FullName != "" {
		user.FullName = updateRequest.FullName
	}

	if updateRequest.Email != "" {
		user.Email = updateRequest.Email
	}

	if updateRequest.Password != "" {
		hashedPassword, errorHandler := utils.PasswordHashing(updateRequest.Password)
		if errorHandler != nil {
			utils.SendErrorResponse(c, "[UPDATE PROFILE] Error hashing password", 500)
			return
		}
		user.HashedPassword = hashedPassword
	}

	if errorHandler := repositories.UpdateUserByUserID(user); errorHandler != nil {
		utils.SendErrorResponse(c, "[UPDATE PROFILE] Error updating user", 500)
		return
	}

	utils.SendSuccessResponse(c, "[UPDATE PROFILE] User profile updated successfully", user)
}

func DeleteUserByUserID(c *gin.Context) {
	userIDStr := c.Param("id")

	if userIDStr == "" {
		utils.SendErrorResponse(c, "[DELETE USER] User ID is required", 400)
		return
	}

	userID, errorHandler := strconv.Atoi(userIDStr)
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[DELETE USER] Invalid user ID", 400)
		return
	}

	if errorHandler := repositories.DeleteExploreChapterByUserID(uint(userID)); errorHandler != nil {
		utils.SendErrorResponse(c, "[DELETE USER] Error deleting user's explore courses", 500)
		return
	}

	if errorHandler := repositories.DeleteExploreCourseByUserID(uint(userID)); errorHandler != nil {
		utils.SendErrorResponse(c, "[DELETE USER] Error deleting user's explore courses", 500)
		return
	}

	if errorHandler := repositories.DeleteCourseByUserID(uint(userID)); errorHandler != nil {
		utils.SendErrorResponse(c, "[DELETE USER] Error deleting user's explore courses", 500)
		return
	}

	if errorHandler := repositories.DeleteChapterByUserID(uint(userID)); errorHandler != nil {
		utils.SendErrorResponse(c, "[DELETE USER] Error deleting user's explore courses", 500)
		return
	}

	if errorHandler := repositories.DeleteUserByUserID(uint(userID)); errorHandler != nil {
		utils.SendErrorResponse(c, "[DELETE USER] Error deleting user", 500)
		return
	}

	utils.SendSuccessResponse(c, "[DELETE USER] User deleted successfully", nil)
}