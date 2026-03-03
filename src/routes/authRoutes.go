package routes

import (
	"segmenta/src/handlers"
	"segmenta/src/middlewares"
	"segmenta/src/models"

	"github.com/gin-gonic/gin"
)

// Auth Routes
func SetupAuthRoutes(router *gin.Engine, config *models.AppConfig) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", handlers.Register)
		authGroup.POST("/login", handlers.Login)
		authGroup.POST("/logout", handlers.Logout)
		authGroup.POST("/refresh", middlewares.AuthMiddleware(), handlers.Refresh)
	}
}