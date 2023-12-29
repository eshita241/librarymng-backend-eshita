package models

import (
	"time"

	"gorm.io/gorm"
)

// Transaction represents a user transaction in the database
type Issue struct {
	gorm.Model
	ID              uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID          uint      `gorm:"not null;uniqueIndex:idx_transaction_user" json:"user_id"`
	User            User      `gorm:"foreignKey:UserID;references:ID" json:"user"`
	BookID          uint      `gorm:"not null;uniqueIndex:idx_transaction_book" json:"book_id"`
	Book            Book      `gorm:"foreignKey:BookID;references:ID" json:"book"`
	TransactionDate time.Time `gorm:"not null" json:"transaction_date"`
	DueDate         time.Time `gorm:"not null" json:"due_date"`
	ReturnDate      time.Time `gorm:"" json:"return_date"`
	Fine            uint      `gorm:"not null" json:"fine"`
	Status          string    `gorm:"size:255;not null;check:status IN ('borrowed', 'returned', 'overdue')" json:"status"`
	Notes           string    `gorm:"size:255;not null;" json:"notes"`
}
