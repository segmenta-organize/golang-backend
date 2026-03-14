package models

import (
	"time"
)

type Course struct {
	CourseID             uint      `gorm:"primaryKey;autoIncrement;column:course_id"`
	UserID               int       `gorm:"not null;column:user_id"`
	Title                string    `gorm:"type:varchar(200);not null;column:title"`
	Description          *string   `gorm:"type:text;column:description"`
	Channel              *string   `gorm:"type:varchar(150);column:channel"`
	ChannelLink          *string   `gorm:"type:text;column:channel_link"`
	VideoLink            *string   `gorm:"type:text;column:video_link"`
	ThumbnailImageURL    *string   `gorm:"type:text;column:thumbnail_image_url"`
	Progress             int       `gorm:"default:0;check:progress >= 0 AND progress <= 100;column:progress"`
	SourcePublicCourseID *int      `gorm:"column:source_public_course_id"`
	SourceVersion        int       `gorm:"not null;column:source_version"`
	UpdateCheck          bool      `gorm:"not null;default:false;column:update_check"`
	CreatedAt            time.Time `gorm:"autoCreateTime;column:created_at"`
	UpdatedAt            time.Time `gorm:"autoUpdateTime;column:updated_at"`
}

type Chapter struct {
	ChapterID      uint      `gorm:"primaryKey;autoIncrement;column:chapter_id"`
	CourseID       int       `gorm:"not null;column:course_id"`
	Title          string    `gorm:"type:varchar(200);not null;column:title"`
	VideoTimestamp *string   `gorm:"type:varchar(50);column:video_timestamp"`
	Position       int       `gorm:"not null;check:position > 0;column:position"`
	IsCompleted    bool      `gorm:"default:false;column:is_completed"`
	CreatedAt      time.Time `gorm:"autoCreateTime;column:created_at"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime;column:updated_at"`
}