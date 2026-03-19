package repositories

import (
	"segmenta/src/configs"
	"segmenta/src/models"
)

func CreateUser(user *models.User) error {
	return configs.Database.Create(user).Error
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := configs.Database.Where("email = ?", email).First(&user)
	return &user, result.Error
}

func GetUserByUserID(id uint) (*models.User, error) {
	var user models.User
	if errorHandler := configs.Database.First(&user, id).Error; errorHandler != nil {
		return nil, errorHandler
	}
	return &user, nil
}

func UpdateUserByUserID(user *models.User) error {
	return configs.Database.Where("id = ?", user.UserID).Updates(user).Error
}

func DeleteUserByUserID(id uint) error {
	return configs.Database.Delete(&models.User{}, id).Error
}