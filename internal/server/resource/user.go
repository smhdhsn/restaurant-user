package resource

import (
	"github.com/smhdhsn/restaurant-user/internal/server/handler"
)

// UserResource holds user resource's handlers.
type UserResource struct {
	SourceHandler *handler.SourceHandler
}

// NewUserResource creates a new user resource with all handlers within itself.
func NewUserResource(source *handler.SourceHandler) *UserResource {
	return &UserResource{
		SourceHandler: source,
	}
}
