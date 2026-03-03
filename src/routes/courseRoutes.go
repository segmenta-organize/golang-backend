package routes

import (
	"segmenta/src/handlers"
	"segmenta/src/middlewares"
	"segmenta/src/models"

	"github.com/gin-gonic/gin"
)

func SetupCourseRoutes(router *gin.Engine, config *models.AppConfig) {
	courseGroup := router.Group("/courses")
	{
		courseGroup.GET("/", middlewares.AuthMiddleware(), handlers.GetCourses)
		courseGroup.GET("/:id", middlewares.AuthMiddleware(), handlers.GetCourseByID)
		courseGroup.POST("/create", middlewares.AuthMiddleware(), handlers.CreateCourse)
		courseGroup.PUT("/edit/:id", middlewares.AuthMiddleware(), handlers.UpdateCourse)
		courseGroup.DELETE("/delete/:id", middlewares.AuthMiddleware(), handlers.DeleteCourse)
	}
}