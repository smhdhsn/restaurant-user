package contract

import (
	"github.com/pkg/errors"

	"github.com/smhdhsn/restaurant-user/internal/repository/entity"
)

// This section contains MySQL error numbers.
const (
	DuplicateEntryErrNum = 1062
)

// This block holds common errors that might happen within repository.
var (
	ErrRecordNotFound = errors.New("record_not_found")
	ErrDuplicateEntry = errors.New("duplicate_entry")
)

// UserRepository is the interface representing user repository or it's mock.
type UserRepository interface {
	FindBy(*entity.User) (*entity.User, error)
	Store(*entity.User) (*entity.User, error)
}
