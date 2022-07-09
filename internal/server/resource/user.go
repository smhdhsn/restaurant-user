package resource

import (
	authProto "github.com/smhdhsn/restaurant-user/internal/protos/user/auth"
)

// UserResource holds user resource's handlers.
type UserResource struct {
	AuthHandler authProto.UserAuthServiceServer
}

// NewUserResource creates a new user resource with all handlers within itself.
func NewUserResource(ah authProto.UserAuthServiceServer) *UserResource {
	return &UserResource{
		AuthHandler: ah,
	}
}
