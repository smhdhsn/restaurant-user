package model

import (
	"gorm.io/gorm"
)

// User represents the users table on database.
type User struct {
	gorm.Model
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Email     string `gorm:"unique,not null"`
	Status    string `gorm:"not null"`
	Password  string `gorm:"not null" json:"-"`
}
