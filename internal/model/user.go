package model

import (
	"time"
)

// User represents the users table on database.
type User struct {
	ID        uint32 `gorm:"primaryKey"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Email     string `gorm:"size:255;uniqueIndex;not null"`
	Status    string `gorm:"type:enum('active', 'inactive');not null"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// UserDTO represents user's data transfer object.
type UserDTO struct {
	ID        uint32
	FirstName string
	LastName  string
	Email     string
	Status    string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
