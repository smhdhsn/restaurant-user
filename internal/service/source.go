package service

import (
	"github.com/smhdhsn/restaurant-user/internal/model"
	"github.com/smhdhsn/restaurant-user/pkg/encryption"

	repositoryContract "github.com/smhdhsn/restaurant-user/internal/repository/contract"
	serviceContract "github.com/smhdhsn/restaurant-user/internal/service/contract"
)

// SourceService contains repositories that will be used within this service.
type SourceService struct {
	uRepo repositoryContract.UserRepository
}

// NewSourceService creates a user's source service with it's dependencies.
func NewSourceService(uRepo repositoryContract.UserRepository) serviceContract.UserSourceService {
	return &SourceService{
		uRepo: uRepo,
	}
}

// Store is responsible for storing user data inside database.
func (s *SourceService) Store(u *model.UserDTO) (*model.UserDTO, error) {
	u.Password = encryption.EncodeMD5(u.Password)

	return s.uRepo.Store(u)
}

// Find is responsible for fetching user's full details from database.
func (s *SourceService) Find(u *model.UserDTO) (*model.UserDTO, error) {
	return s.uRepo.Find(u)
}

// Destroy is responsible for deleting a user from the database.
func (s *SourceService) Destroy(u *model.UserDTO) error {
	return s.uRepo.Destroy(u)
}

// Update is responsible for updating user's information inside database.
func (s *SourceService) Update(u *model.UserDTO) error {
	if u.Password != "" {
		u.Password = encryption.EncodeMD5(u.Password)
	}

	return s.uRepo.Update(u)
}
