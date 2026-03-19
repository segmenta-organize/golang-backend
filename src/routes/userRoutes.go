package routes

import (
	"segmenta/src/handlers"
	"segmenta/src/middlewares"
	"segmenta/src/models"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine, config *models.AppConfig) {
	userGroup := router.Group("/user")
	{
		userGroup.GET("/:id", middlewares.AuthMiddleware(), handlers.GetOneUserByUserID)
		userGroup.PUT("/:id/edit", middlewares.AuthMiddleware(), handlers.UpdateUserByUserID)
		userGroup.DELETE("/:id/delete", middlewares.AuthMiddleware(), handlers.DeleteUserByUserID)
	}
}