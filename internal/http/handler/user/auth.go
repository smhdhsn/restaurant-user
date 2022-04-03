package user

import (
	uService "github.com/smhdhsn/bookstore-user/internal/service/user"
)

// Auth contains services that can be used within user auth handler.
type Auth struct {
	authService *uService.AuthService
}

// NewAuth creates a new user auth handler.
func NewAuth(authService *uService.AuthService) *Auth {
	return &Auth{
		authService: authService,
	}
}
