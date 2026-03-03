package utils

import (
	"github.com/gin-gonic/gin"
)

func LinkDuplicateCheck(c *gin.Context, prefix string, fieldName string, value string) bool {
	userID, ok := getUserID(c, prefix)
	if !ok {
		return false
	}

	courses, errorHandler := repositories.GetAllCourses(userID)
	if errorHandler != nil {
		SendErrorResponse(c, prefix+" Error fetching courses", 500)
		return false
	}

	for _, course := range courses {
		if fieldName == "course_link" && course.CourseLink == value {
			SendErrorResponse(c, prefix+" Course link already exists", 400)
			return true
		}
	}

	return false
}