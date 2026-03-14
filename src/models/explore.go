package models

import (
	"time"
)

type ExploreCourse struct {
	ExploreCourseCourseID uint      `gorm:"primaryKey;autoIncrement;column:explore_course_id"`
	CreatorID             int       `gorm:"not null;column:creator_id"`
	Title                 string    `gorm:"type:varchar(200);not null;column:title"`
	Description           *string   `gorm:"type:text;column:description"`
	Channel               *string   `gorm:"type:varchar(150);column:channel"`
	ChannelLink           *string   `gorm:"type:text;column:channel_link"`
	VideoLink             *string   `gorm:"type:text;column:video_link"`
	ThumbnailImageURL     *string   `gorm:"type:text;column:thumbnail_image_url"`
	Version        		  *int      `gorm:"not null;column:version"`
	CreatedAt             time.Time `gorm:"autoCreateTime;column:created_at"`
	UpdatedAt             time.Time `gorm:"autoUpdateTime;column:updated_at"`
}

type ExploreChapter struct {
	ExploreChapterID      uint      `gorm:"primaryKey;autoIncrement;column:explore_chapter_id"`
	ExploreCourseID       uint      `gorm:"not null;column:explore_course_id"`
	Title                 string    `gorm:"type:varchar(200);not null;column:title"`
	Description           *string   `gorm:"type:text;column:description"`
	Order                 int       `gorm:"not null;column:order"`
	CreatedAt             time.Time `gorm:"autoCreateTime;column:created_at"`
	UpdatedAt             time.Time `gorm:"autoUpdateTime;column:updated_at"`
}