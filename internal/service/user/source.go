package user

import (
	"github.com/smhdhsn/bookstore-user/internal/model"
	"github.com/smhdhsn/bookstore-user/internal/repository/contract"
	"github.com/smhdhsn/bookstore-user/internal/request"
	"github.com/smhdhsn/bookstore-user/util/encrypt"
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

// Find is responsible for storing user data inside database.
func (s *SourceService) Find(userID uint) (*model.User, error) {
	return s.uRepo.Find(userID)
}

// Store is responsible for storing user data inside database.
func (s *SourceService) Store(req *request.StoreUserReq) (user *model.User, err error) {
	user = &model.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  encrypt.ToMD5(req.Password),
		Status:    req.Status,
	}

	if err := s.uRepo.Store(user); err != nil {
		return nil, err
	}
	return
}

// Update is responsible for updating user's information inside database.
func (s *SourceService) Update(req *request.UpdateUserReq, userID uint) error {
	user := model.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
	}

	return s.uRepo.Update(&user, userID)
}

// Destroy is responsible for deleting a user from the database.
func (s *SourceService) Destroy(userID uint) error {
	return s.uRepo.Destroy(userID)
}
