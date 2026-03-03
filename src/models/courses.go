package models

import (
	"time"
)

type Course struct {
	CourseID            uint      `gorm:"primaryKey;autoIncrement;column:course_id"`
	UserID              int       `gorm:"not null;column:user_id"`
	Title               string    `gorm:"type:varchar(200);not null;column:title"`
	Description         *string   `gorm:"type:text;column:description"`
	Channel             *string   `gorm:"type:varchar(150);column:channel"`
	Link                *string   `gorm:"type:text;column:link"`
	ImageURL            *string   `gorm:"type:text;column:image_url"`
	Progress            int       `gorm:"default:0;check:progress >= 0 AND progress <= 100;column:progress"`
	SourcePublicCourseID *int     `gorm:"column:source_public_course_id"`
	SourceVersion       *int      `gorm:"column:source_version"`
	CreatedAt           time.Time `gorm:"autoCreateTime;column:created_at"`
	UpdatedAt           time.Time `gorm:"autoUpdateTime;column:updated_at"`
}