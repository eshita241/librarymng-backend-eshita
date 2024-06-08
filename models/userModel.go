{/*package models

import (
	"gorm.io/gorm"
)

// User is a representation of a user in the database
type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string `gorm:"size:255;not null;" json:"name"`
	Email    string `gorm:"size:255;not null;uniqueIndex" json:"email"`
	Password string `gorm:"size:100;not null" json:"password"`
	Role     string `gorm:"size:50;not null" json:"role"` // Ensure this matches the database schema
}
*/}