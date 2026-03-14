package utils

import (
	"segmenta/src/repositories"

	"github.com/gin-gonic/gin"
)

func LinkDuplicateCheck(c *gin.Context, prefix string, fieldName string, value string) bool {
	userID, ok := GetUserID(c, prefix)
	if !ok {
		return false
	}

	// Cek duplikat langsung ke database
	exists, errorHandler := repositories.CheckCourseLinkExists(userID, value)
	if errorHandler != nil {
		SendErrorResponse(c, prefix+" Error checking course link", 500)
		return false
	}
	if exists {
		SendErrorResponse(c, prefix+" Course link already exists", 400)
		return true
	}

	return false
}
