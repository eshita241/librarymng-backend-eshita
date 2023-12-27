package models

import (
	"gorm.io/gorm"
)

// User is a representation of a user in the database
type Book struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Title     string `gorm:"size:255;not null;" json:"title"`
	Author    string `gorm:"size:255;not null" json:"author"`
	Category  string `gorm:"size:255;not null" json:"category"`
	Publisher string `gorm:"size:255;not null" json:"publisher"`
	Year      uint   `gorm:"size:255;not null" json:"year"`
	Edition   string `gorm:"size:255;not null" json:"edition"`
	Language  string `gorm:"size:255;not null" json:"language"`
	Copies    uint   `gorm:"size:255;not null" json:"copies"`
	Condition string `gorm:"size:255;not null" json:"condition"`
}
