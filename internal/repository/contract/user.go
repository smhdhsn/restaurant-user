package contract

import (
	"errors"

	"github.com/smhdhsn/restaurant-user/internal/model"
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
	Store(*model.UserDTO) (*model.UserDTO, error)
	Find(*model.UserDTO) (*model.UserDTO, error)
	Destroy(*model.UserDTO) error
	Update(*model.UserDTO) error
}
