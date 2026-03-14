package handlers

import (
	"segmenta/src/controllers"
	
	"github.com/gin-gonic/gin"
)

// Explore Course Handlers

func GetAllCoursesForExplore(c *gin.Context) {
	controllers.GetAllCoursesForExplore(c)
}

func GetExploredCourseByID(c *gin.Context) {
	controllers.GetExploredCourseByID(c)
}

func SearchCourses(c *gin.Context) {
	controllers.SearchCourses(c)
}

func EnrollInCourse(c *gin.Context) {
	controllers.EnrollInCourse(c)
}

// Explore Course Chapter Handlers

func GetAllExploreChapterByCourseID(c *gin.Context) {
	controllers.GetAllExploreChapterByCourseID(c)
}

func GetOneExploreChapterByID(c *gin.Context) {
	controllers.GetOneExploreChapterByID(c)
}

func CreateExploreChapter(c *gin.Context) {
	controllers.CreateExploreChapter(c)
}

func UpdateExploreChapter(c *gin.Context) {
	controllers.UpdateExploreChapter(c)
}

func DeleteExploreChapter(c *gin.Context) {
	controllers.DeleteExploreChapter(c)
}