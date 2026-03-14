package handlers

import (
	"segmenta/src/controllers"

	"github.com/gin-gonic/gin"
)

// Course Handlers

func GetCourses(c *gin.Context) {
	controllers.GetAllCourses(c)
}

func GetCourseByID(c *gin.Context) {
	controllers.GetCourseByID(c)
}

func CreateCourse(c *gin.Context) {
	controllers.CreateCourse(c)
}

func AutoCreateCourse(c *gin.Context) {
	controllers.AutoCreateCourses(c)
}

func UpdateCourse(c *gin.Context) {
	controllers.UpdateCourse(c)
}

func DeleteCourse(c *gin.Context) {
	controllers.DeleteCourse(c)
}

// Chapter Handlers

func GetAllChaptersByCourseID(c *gin.Context) {
	controllers.GetAllChaptersByCourseID(c)
}

func GetOneChapterByID(c *gin.Context) {
	controllers.GetOneChapterByID(c)
}

func CreateChapter(c *gin.Context) {
	controllers.CreateChapter(c)
}

func UpdateChapter(c *gin.Context) {
	controllers.UpdateChapter(c)
}

func DeleteChapter(c *gin.Context) {
	controllers.DeleteChapter(c)
}
