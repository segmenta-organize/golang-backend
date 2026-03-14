package routes

import (
	"segmenta/src/models"

	"github.com/gin-gonic/gin"
)

func SetupAllRoutes(router *gin.Engine, config *models.AppConfig) {
	SetupAuthRoutes(router, config)
	SetupCourseRoutes(router, config)
	SetupChapterRoutes(router, config)
	SetupExploreRoutes(router, config)
	SetupExploreChapterRoutes(router, config)
}