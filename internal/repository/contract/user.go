package contract

import (
	"github.com/smhdhsn/bookstore-user/internal/model"
)

// UserRepository is the interface representing user repository or it's mock.
type UserRepository interface {
	Store(*model.User) (*model.User, error)
	Update(*model.User, uint) error
	Find(uint) (*model.User, error)
	Destroy(uint) error
}
