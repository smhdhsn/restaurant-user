package resource

import (
	uspb "github.com/smhdhsn/restaurant-user/internal/protos/user/source"
)

// UserResource holds user resource's handlers.
type UserResource struct {
	SourceHandler uspb.UserSourceServiceServer
}

// NewUserResource creates a new user resource with all handlers within itself.
func NewUserResource(source uspb.UserSourceServiceServer) *UserResource {
	return &UserResource{
		SourceHandler: source,
	}
}
