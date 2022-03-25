package model

import (
	"gorm.io/gorm"
)

// User represents the users table on database.
type User struct {
	gorm.Model
	FirstName string `gorm:"not null" validate:"required"`
	LastName  string `gorm:"not null" validate:"required"`
	Email     string `gorm:"uniqueIndex,not null" validate:"required"`
	Status    string `gorm:"not null" validate:"required"`
	Password  string `gorm:"not null" validate:"required"`
}
