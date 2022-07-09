package entity

import (
	"time"
)

// User represents the user repository's entity.
type User struct {
	ID        uint32
	FirstName string
	LastName  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
