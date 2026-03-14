package repositories

import (
	"segmenta/src/configs"
	"segmenta/src/models"
)

// Course Repositories

func GetAllCourses(userID uint) ([]models.Course, error) {
	var courses []models.Course
	result := configs.Database.Where("user_id = ?", userID).Find(&courses)
	return courses, result.Error
}

func GetCourseByID(id uint) (*models.Course, error) {
	var course models.Course
	result := configs.Database.First(&course, "course_id = ?", id)
	return &course, result.Error
}

func CreateCourse(course *models.Course) error {
	return configs.Database.Create(course).Error
}

func UpdateCourseByID(id uint, course *models.Course) error {
	return configs.Database.Model(&models.Course{}).Where("course_id = ?", id).Updates(course).Error
}

func DeleteCourseByID(id uint) error {
	return configs.Database.Where("course_id = ?", id).Delete(&models.Course{}).Error
}

func CheckCourseLinkExists(userID uint, link string) (bool, error) {
	var count int64
	result := configs.Database.Model(&models.Course{}).Where("user_id = ? AND video_link = ?", userID, link).Count(&count)
	return count > 0, result.Error
}

// Chapter Repositories

func GetAllChaptersByCourseID(courseID uint) ([]models.Chapter, error) {
	var chapters []models.Chapter
	result := configs.Database.Where("course_id = ?", courseID).Order("position ASC").Find(&chapters)
	return chapters, result.Error
}

func GetChapterByID(id uint) (*models.Chapter, error) {
	var chapter models.Chapter
	result := configs.Database.First(&chapter, "chapter_id = ?", id)
	return &chapter, result.Error
}

func CreateChapter(chapter *models.Chapter) error {
	return configs.Database.Create(chapter).Error
}

func UpdateChapterByID(id uint, chapter *models.Chapter) error {
	return configs.Database.Model(&models.Chapter{}).Where("chapter_id = ?", id).Updates(chapter).Error
}

func DeleteChapterByID(id uint) error {
	return configs.Database.Where("chapter_id = ?", id).Delete(&models.Chapter{}).Error
}