package models

import (
	"time"
)

type Chapter struct {
	ChapterID      int       `json:"chapter_id" gorm:"primaryKey;autoIncrement"`
	CourseID       int       `json:"course_id" gorm:"not null"`
	Title          string    `json:"title" gorm:"type:varchar(200);not null"`
	VideoTimestamp string    `json:"video_timestamp" gorm:"type:varchar(50)"`
	Position       int       `json:"position" gorm:"not null"`
	IsCompleted    bool      `json:"is_completed" gorm:"default:false"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Course Course `json:"course" gorm:"foreignKey:CourseID;references:CourseID;constraint:OnDelete:CASCADE"`
}