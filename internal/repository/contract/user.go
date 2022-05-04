package contract

import (
	"github.com/smhdhsn/bookstore-user/internal/model"
)

// UserRepository is the interface representing user repository or it's mock.
type UserRepository interface {
	Store(*model.UserDTO) (*model.UserDTO, error)
	FindBy(FilterBy) (model.UserDTOList, error)
	Find(uint) (*model.UserDTO, error)
	Show(uint) (*model.UserDTO, error)
	Update(*model.UserDTO, uint) error
	Destroy(uint) error
}
