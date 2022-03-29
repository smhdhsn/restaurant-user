package model

import (
	"time"
)

// User represents the users table on database.
type User struct {
	ID        uint   `gorm:"primaryKey"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Email     string `gorm:"size:255;uniqueIndex;not null"`
	Status    string `gorm:"type:enum('active', 'inactive');not null"`
	Password  string `gorm:"not null" json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
