package resource

import (
	uHand "github.com/smhdhsn/restaurant-user/internal/http/handler/user"
)

// UserResource holds user resource's handlers.
type UserResource struct {
	SourceHandler *uHand.SourceHandler
}

// NewUserResource creates a new user resource with all handlers within itself.
func NewUserResource(source *uHand.SourceHandler) *UserResource {
	return &UserResource{
		SourceHandler: source,
	}
}
