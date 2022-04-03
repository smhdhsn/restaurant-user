package resource

import (
	"github.com/smhdhsn/bookstore-user/internal/http/handler/user"
)

// UserResource holds user resource's handlers.
type UserResource struct {
	Source *user.Source
	Search *user.Search
	Auth   *user.Auth
}

// NewUserResource creates a new user resource with all handlers within itself.
func NewUserResource(source *user.Source, search *user.Search, auth *user.Auth) *UserResource {
	return &UserResource{
		Source: source,
		Search: search,
		Auth:   auth,
	}
}
