package user

import (
	"github.com/smhdhsn/restaurant-user/internal/repository/contract"
)

// AuthService contains repositories that will be used within this service.
type AuthService struct {
	uRepo contract.UserRepository
}

// NewAuthService creates a user's auth service with it's dependencies.
func NewAuthService(uRepo contract.UserRepository) *AuthService {
	return &AuthService{
		uRepo: uRepo,
	}
}
