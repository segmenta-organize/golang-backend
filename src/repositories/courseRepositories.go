package repositories

import (
	"segmenta/src/configs"
	"segmenta/src/models"
)

func GetAllCourses(userID uint) ([]models.Course, error) {
	var courses []models.Course
	result := configs.Database.Where("user_id = ?", userID).Find(&courses)
	return courses, result.Error
}

func GetCourseByID(id uint) (*models.Course, error) {
	var course models.Course
	result := configs.Database.First(&course, id)
	return &course, result.Error
}

func CreateCourse(course *models.Course) error {
	return configs.Database.Create(course).Error
}

func UpdateCourseByID(id uint, course *models.Course) error {
	return configs.Database.Model(&models.Course{}).Where("id = ?", id).Updates(course).Error
}

func DeleteCourseByID(id uint) error {
	return configs.Database.Delete(&models.Course{}, id).Error
}