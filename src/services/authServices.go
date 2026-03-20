package services

import (
	"segmenta/src/models"
	"segmenta/src/repositories"
	"segmenta/src/utils"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var RegisterRequest struct {
		FullName string `json:"full_name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if errorHandler := c.ShouldBindJSON(&RegisterRequest); errorHandler != nil {
		utils.SendErrorResponse(c, "[REGISTER] Invalid request data", 400)
		return
	}

	if _, errorHandler := repositories.GetUserByEmail(RegisterRequest.Email); errorHandler == nil {
		utils.SendErrorResponse(c, "[REGISTER] Email already registered", 400)
		return
	}

	hashedPassword, errorHandler := utils.PasswordHashing(RegisterRequest.Password)
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[REGISTER] Error hashing password", 500)
		return
	}

	user := models.User{
		FullName:       RegisterRequest.FullName,
		Email:          RegisterRequest.Email,
		HashedPassword: hashedPassword,
	}
	if errorHandler := repositories.CreateUser(&user); errorHandler != nil {
		utils.SendErrorResponse(c, "[REGISTER] Error creating user", 500)
		return
	}

	token, errorHandler := utils.GenerateJWT(&user)
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[REGISTER] Error generating token but account created", 500)
		return
	}

	utils.SendSuccessResponse(c, "[REGISTER] User registered successfully", gin.H{"token": token})
}

func Login(c *gin.Context) {
	var LoginRequest struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if errorHandler := c.ShouldBindJSON(&LoginRequest); errorHandler != nil {
		utils.SendErrorResponse(c, "[LOGIN] Invalid request data", 400)
		return
	}

	user, errorHandler := repositories.GetUserByEmail(LoginRequest.Email)
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[LOGIN] User not found", 404)
		return
	}

	if errorHandler := utils.ComparePasswords(user.HashedPassword, LoginRequest.Password); errorHandler != nil {
		utils.SendErrorResponse(c, "[LOGIN] Incorrect password", 401)
		return
	}

	token, errorHandler := utils.GenerateJWT(user)
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[LOGIN] Error generating token", 500)
		return
	}

	utils.SendSuccessResponse(c, "[LOGIN] Login successful", gin.H{"token": token})
}

func Logout(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		utils.SendErrorResponse(c, "[LOGOUT] Authorization token required", 401)
		return
	}

	utils.SendSuccessResponse(c, "[LOGOUT] Logout successful", nil)
}

func Refresh(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		utils.SendErrorResponse(c, "[REFRESH] User ID not found in context", 500)
		return
	}

	var id uint
	switch v := userID.(type) {

	case float64:
		id = uint(v)
	case uint:
		id = v
	default:
		utils.SendErrorResponse(c, "[REFRESH] Invalid user ID type", 500)
		return
	}

	user, errorHandler := repositories.GetUserByUserID(id)
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[REFRESH] User not found", 404)
		return
	}

	newToken, errorHandler := utils.GenerateJWT(user)
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[REFRESH] Error generating new token", 500)
		return
	}

	utils.SendSuccessResponse(c, "[REFRESH] Token refreshed successfully", gin.H{"token": newToken})
}

func ForgotPassword(c *gin.Context) {
	var forgotPasswordRequest struct {
		Email string `json:"email" binding:"required,email"`
	}

	if errorHandler := c.ShouldBindJSON(&forgotPasswordRequest); errorHandler != nil {
		utils.SendErrorResponse(c, "[FORGOT PASSWORD] Invalid request data", 400)
		return
	}

	user, errorHandler := repositories.GetUserByEmail(forgotPasswordRequest.Email)
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[FORGOT PASSWORD] User not found", 404)
		return
	}

	resetLink := utils.GenerateResetPasswordLink(user.Email)
	if resetLink == "" {
		utils.SendErrorResponse(c, "[FORGOT PASSWORD] Error generating reset link", 500)
		return
	}

	errorHandler = utils.SendResetPasswordtoEmail(user.Email, resetLink)
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[FORGOT PASSWORD] Error sending reset password email", 500)
		return
	}

	utils.SendSuccessResponse(c, "[FORGOT PASSWORD] Forgot password request successful", gin.H{"status": "success"})
}

func ResetPassword(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		utils.SendErrorResponse(c, "[RESET PASSWORD] Token is required", 400)
		return
	}

	var resetPasswordRequest struct {
		NewPassword string `json:"new_password" binding:"required,min=6"`
	}

	if errorHandler := c.ShouldBindJSON(&resetPasswordRequest); errorHandler != nil {
		utils.SendErrorResponse(c, "[RESET PASSWORD] Invalid request data", 400)
		return
	}

	email, errorHandler := utils.ParseResetPasswordToken(token)
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[RESET PASSWORD] Invalid or expired token", 400)
		return
	}

	user, errorHandler := repositories.GetUserByEmail(email)
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[RESET PASSWORD] User not found", 404)
		return
	}

	hashedPassword, errorHandler := utils.PasswordHashing(resetPasswordRequest.NewPassword)
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[RESET PASSWORD] Error hashing password", 500)
		return
	}

	user.HashedPassword = hashedPassword
	if errorHandler := repositories.UpdateUserByUserID(user); errorHandler != nil {
		utils.SendErrorResponse(c, "[RESET PASSWORD] Error updating password", 500)
		return
	}

	utils.SendSuccessResponse(c, "[RESET PASSWORD] Password reset successful", nil)
}