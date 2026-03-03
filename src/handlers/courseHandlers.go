package handlers

import (
	"segmenta/src/controllers"

	"github.com/gin-gonic/gin"
)

func GetCourses(c *gin.Context) {
	controllers.GetAllCourses(c)
}

func GetCourseByID(c *gin.Context) {
	controllers.GetCourseByID(c)
}

func CreateCourse(c *gin.Context) {
	controllers.CreateCourse(c)
}

func UpdateCourse(c *gin.Context) {
	controllers.UpdateCourse(c)
}

func DeleteCourse(c *gin.Context) {
	controllers.DeleteCourse(c)
}