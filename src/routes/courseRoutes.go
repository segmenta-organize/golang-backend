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
		courseGroup.GET("/", middlewares.AuthMiddleware(), handlers.GetAllEnrolledCourses)
		courseGroup.GET("/:id", middlewares.AuthMiddleware(), handlers.GetOneCourseWithChaptersByID)
		courseGroup.POST("/create", middlewares.AuthMiddleware(), handlers.CreateManualCourseWithChapters)
		courseGroup.POST("/auto-create", middlewares.AuthMiddleware(), handlers.AutoCreateCourseWithChapters)
		courseGroup.PUT("/:id/edit", middlewares.AuthMiddleware(), handlers.UpdateCourseWithChapters)
		courseGroup.PUT("/:id/auto-edit", middlewares.AuthMiddleware(), handlers.AutoUpdateCourseWithChapters)
		courseGroup.DELETE("/:id/delete", middlewares.AuthMiddleware(), handlers.DeleteOneCourseByID)
		courseGroup.POST("/:id/create-public", middlewares.AuthMiddleware(), handlers.CreatePublicCourseFromCourse)
		courseGroup.PUT("/:id/update-public", middlewares.AuthMiddleware(), handlers.UpdatePublicCourseFromCourse)
	}
}

func SetupChapterRoutes(router *gin.Engine, config *models.AppConfig) {
	chapterGroup := router.Group("/chapters")
	{
		chapterGroup.GET("/", middlewares.AuthMiddleware(), handlers.GetAllChaptersByCourseID)
		chapterGroup.GET("/:id", middlewares.AuthMiddleware(), handlers.GetOneChapterByID)
		chapterGroup.POST("/create", middlewares.AuthMiddleware(), handlers.CreateChapter)
		chapterGroup.PUT("/:id/edit", middlewares.AuthMiddleware(), handlers.UpdateChapter)
		chapterGroup.DELETE("/:id/delete", middlewares.AuthMiddleware(), handlers.DeleteChapter)
		chapterGroup.POST("/:id/create-public", middlewares.AuthMiddleware(), handlers.CreatePublicChapterFromChapter)
		chapterGroup.PUT("/:id/update-public", middlewares.AuthMiddleware(), handlers.UpdatePublicChapterFromChapter)
		chapterGroup.DELETE("/:id/delete-public", middlewares.AuthMiddleware(), handlers.DeletePublicChapterFromChapter)
	}
}