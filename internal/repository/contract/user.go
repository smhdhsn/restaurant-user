package contract

import (
	"github.com/smhdhsn/bookstore-user/internal/model"
)

// UserRepository is the interface representing user repository or it's mock.
type UserRepository interface {
	Store(*model.User) (*model.User, error)
	FindBy(FilterBy) ([]*model.User, error)
	Find(uint) (*model.User, error)
	Update(*model.User, uint) error
	Destroy(uint) error
}
