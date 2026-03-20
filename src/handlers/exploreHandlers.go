package handlers

import (
	"segmenta/src/services"

	"github.com/gin-gonic/gin"
)

// Explore Course Handlers

func GetAllExploreCourses(c *gin.Context) {
	services.GetAllExploreCourses(c)
}

func GetExploreCourseByID(c *gin.Context) {
	services.GetExploreCourseByID(c)
}

func SearchCourses(c *gin.Context) {
	services.SearchCourses(c)
}

func GetAllCoursesByCategoryForExplore(c *gin.Context) {
	services.GetAllCoursesByCategoryForExplore(c)
}

func EnrollInCourse(c *gin.Context) {
	services.EnrollInCourse(c)
}

func EditPublicCourse(c *gin.Context) {
	services.EditPublicCourse(c)
}

func DeletePublicCourse(c *gin.Context) {
	services.DeletePublicCourse(c)
}

// Explore Course Chapter Handlers

func GetAllExploreChapterByExploreCourseID(c *gin.Context) {
	services.GetAllExploreChapterByExploreCourseID(c)
}

func GetOneExploreChapterByCourseID(c *gin.Context) {
	services.GetOneExploreChapterByCourseID(c)
}

func CreateExploreChapter(c *gin.Context) {
	services.CreateExploreChapter(c)
}

func UpdateExploreChapter(c *gin.Context) {
	services.UpdateExploreChapter(c)
}

func DeleteExploreChapter(c *gin.Context) {
	services.DeleteExploreChapter(c)
}
