package routes

import (
	"github.com/gin-gonic/gin"
	"segmenta/src/models"
)

func SetupAllRoutes(router *gin.Engine, config *models.AppConfig) {
	SetupAuthRoutes(router, config)
	SetupCourseRoutes(router, config)
}