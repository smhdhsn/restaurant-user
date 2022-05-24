package contract

import (
	"github.com/smhdhsn/restaurant-user/internal/model"
)

// UserRepository is the interface representing user repository or it's mock.
type UserRepository interface {
	Store(*model.UserDTO) (*model.UserDTO, error)
	Find(*model.UserDTO) (*model.UserDTO, error)
	Destroy(*model.UserDTO) error
	Update(*model.UserDTO) error
}
