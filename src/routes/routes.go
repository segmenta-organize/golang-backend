package routes

import (
	"segmenta/src/models"

	"github.com/gin-gonic/gin"
)

func SetupAllRoutes(router *gin.Engine, config *models.AppConfig) {
	api := router.Group("/api")
	{
		SetupAuthRoutes(api, config)
		SetupCourseRoutes(api, config)
		SetupChapterRoutes(api, config)
		SetupExploreRoutes(api, config)
		SetupExploreChapterRoutes(api, config)
		SetupUserRoutes(api, config) // jangan lupa ini kalau ada
	}
}