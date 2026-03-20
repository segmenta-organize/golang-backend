package handlers

import (
	"segmenta/src/services"

	"github.com/gin-gonic/gin"
)

// Course Handlers

func GetAllEnrolledCourses(c *gin.Context) {
	services.GetAllEnrolledCourses(c)
}

func GetOneCourseWithChaptersByID(c *gin.Context) {
	services.GetOneCourseWithChaptersByID(c)
}

func CreateManualCourseWithChapters(c *gin.Context) {
	services.CreateManualCourseWithChapters(c)
}

func AutoCreateCourseWithChapters(c *gin.Context) {
	services.AutoCreateCourseWithChapters(c)
}

func UpdateCourseWithChapters(c *gin.Context) {
	services.UpdateCourseWithChapters(c)
}

func AutoUpdateCourseWithChapters(c *gin.Context) {
	services.AutoUpdateCourseWithChapters(c)
}

func DeleteOneCourseByID(c *gin.Context) {
	services.DeleteOneCourseByID(c)
}

func CreatePublicCourseFromCourse(c *gin.Context) {
	services.CreatePublicCourseFromCourse(c)
}

func UpdatePublicCourseFromCourse(c *gin.Context) {
	services.UpdatePublicCourseFromCourse(c)
}

// Chapter Handlers

func GetAllChaptersByCourseID(c *gin.Context) {
	services.GetAllChaptersByCourseID(c)
}

func GetOneChapterByID(c *gin.Context) {
	services.GetOneChapterByID(c)
}

func CreateChapter(c *gin.Context) {
	services.CreateChapter(c)
}

func UpdateChapter(c *gin.Context) {
	services.UpdateChapter(c)
}

func DeleteChapter(c *gin.Context) {
	services.DeleteChapter(c)
}

// Chapter → Explore Chapter public handlers

func CreatePublicChapterFromChapter(c *gin.Context) {
	services.CreatePublicChapterFromChapter(c)
}

func UpdatePublicChapterFromChapter(c *gin.Context) {
	services.UpdatePublicChapterFromChapter(c)
}

func DeletePublicChapterFromChapter(c *gin.Context) {
	services.DeletePublicChapterFromChapter(c)
}
