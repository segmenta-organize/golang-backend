package routes

import (
	"segmenta/src/handlers"
	"segmenta/src/middlewares"
	"segmenta/src/models"

	"github.com/gin-gonic/gin"
)

func SetupExploreRoutes(router *gin.RouterGroup, config *models.AppConfig) {
	exploreGroup := router.Group("/explore/courses")
	{
		exploreGroup.GET("/", handlers.GetAllExploreCourses)
		exploreGroup.GET("/:id", handlers.GetExploreCourseByID)
		exploreGroup.GET("/search", handlers.SearchCourses)
		exploreGroup.GET("/categories/:category", handlers.GetAllCoursesByCategoryForExplore)
		exploreGroup.POST("/:id/enroll", middlewares.AuthMiddleware(), handlers.EnrollInCourse)
		exploreGroup.PUT("/:id/edit", middlewares.AuthMiddleware(), handlers.EditPublicCourse)
		exploreGroup.DELETE("/:id/delete", middlewares.AuthMiddleware(), handlers.DeletePublicCourse)
	}
}

func SetupExploreChapterRoutes(router *gin.RouterGroup, config *models.AppConfig) {
	exploreChapterGroup := router.Group("/explore/chapters")
	{
		exploreChapterGroup.GET("/", handlers.GetAllExploreChapterByExploreCourseID)
		exploreChapterGroup.GET("/:id", handlers.GetOneExploreChapterByCourseID)
		exploreChapterGroup.POST("/create", middlewares.AuthMiddleware(), handlers.CreateExploreChapter)
		exploreChapterGroup.PUT("/:id/edit", middlewares.AuthMiddleware(), handlers.UpdateExploreChapter)
		exploreChapterGroup.DELETE("/:id/delete", middlewares.AuthMiddleware(), handlers.DeleteExploreChapter)
	}
}
