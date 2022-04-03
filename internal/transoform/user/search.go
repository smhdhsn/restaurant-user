package user

import (
	"time"
)

// SearchList holds response schema for search list end-point.
type SearchList struct {
	ID        uint      `json:"id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
