package routes

import (
	"segmenta/src/handlers"
	"segmenta/src/middlewares"
	"segmenta/src/models"

	"github.com/gin-gonic/gin"
)

func SetupExploreRoutes(router *gin.Engine, config *models.AppConfig) {
	exploreGroup := router.Group("/explore")
	exploreGroup.Use(middlewares.AuthMiddleware())
	{
		exploreGroup.GET("/courses", handlers.GetAllCoursesForExplore)
		exploreGroup.GET("/courses/:id", handlers.GetExploredCourseByID)
		exploreGroup.GET("/courses/search", handlers.SearchCourses)
		exploreGroup.POST("/courses/:id/enroll", handlers.EnrollInCourse)
	}
}

func SetupExploreChapterRoutes(router *gin.Engine, config *models.AppConfig) {
	exploreChapterGroup := router.Group("/explore/chapters")
	exploreChapterGroup.Use(middlewares.AuthMiddleware())
	{
		exploreChapterGroup.GET("/", handlers.GetAllExploreChapterByCourseID)
		exploreChapterGroup.GET("/:id", handlers.GetOneExploreChapterByID)
		exploreChapterGroup.POST("/create", handlers.CreateExploreChapter)
		exploreChapterGroup.PUT("/edit/:id", handlers.UpdateExploreChapter)
		exploreChapterGroup.DELETE("/delete/:id", handlers.DeleteExploreChapter)
	}
}
