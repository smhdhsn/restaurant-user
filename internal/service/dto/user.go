package dto

import (
	"time"
)

// User represents user's data transfer object.
type User struct {
	ID        uint32
	FirstName string
	LastName  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
