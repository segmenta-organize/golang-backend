package models

import (
	"time"
)

type User struct {
	UserID         uint      `gorm:"primaryKey;autoIncrement;column:user_id" json:"user_id"`
	FullName       string    `gorm:"type:varchar(150);not null;column:full_name" json:"full_name"`
	Email          string    `gorm:"type:varchar(150);uniqueIndex;not null;column:email" json:"email"`
	HashedPassword string    `gorm:"type:text;not null;column:hashed_password" json:"hashed_password"`
	Bio            string    `gorm:"column:bio;type:text"`
	CreatedAt      time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	ProfilePicLink string    `gorm:"column:profile_pic_link;type:text" json:"profile_pic_link"`
}