package contract

import (
	"github.com/pkg/errors"

	"github.com/smhdhsn/restaurant-user/internal/service/dto"
)

// This block holds common errors that might happen within repository.
var (
	ErrRecordNotFound = errors.New("record_not_found")
	ErrDuplicateEntry = errors.New("duplicate_entry")
)

// UserAuthService is the interface that user's auth service must implement.
type UserAuthService interface {
	FindBy(*dto.User) (*dto.User, error)
	Store(*dto.User) (*dto.User, error)
}
