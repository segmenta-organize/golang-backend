package models

type Category struct {
	CategoryID   uint   `gorm:"primaryKey;autoIncrement;column:category_id" json:"category_id"`
	CategoryName string `gorm:"type:varchar(100);not null;uniqueIndex;column:category_name" json:"category_name"`
}	