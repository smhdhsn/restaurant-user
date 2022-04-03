package user

import (
	"github.com/smhdhsn/bookstore-user/internal/model"
	"github.com/smhdhsn/bookstore-user/internal/repository/contract"
	"github.com/smhdhsn/bookstore-user/util/encrypt"

	uRequest "github.com/smhdhsn/bookstore-user/internal/request/user"
)

// SourceService contains repositories that will be used within this service.
type SourceService struct {
	uRepo contract.UserRepository
}

// NewSourceService creates a user's source service with it's dependencies.
func NewSourceService(uRepo contract.UserRepository) *SourceService {
	return &SourceService{
		uRepo: uRepo,
	}
}

// Store is responsible for storing user data inside database.
func (s *SourceService) Store(req uRequest.SourceStoreReq) (*model.User, error) {
	user := &model.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  encrypt.ToMD5(req.Password),
		Status:    req.Status,
	}

	return s.uRepo.Store(user)
}

// Find is responsible for storing user data inside database.
func (s *SourceService) Find(userID uint) (*model.User, error) {
	return s.uRepo.Find(userID)
}

// Destroy is responsible for deleting a user from the database.
func (s *SourceService) Destroy(userID uint) error {
	return s.uRepo.Destroy(userID)
}

// Update is responsible for updating user's information inside database.
func (s *SourceService) Update(req uRequest.SourceUpdateReq, userID uint) error {
	user := model.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
	}

	return s.uRepo.Update(&user, userID)
}
