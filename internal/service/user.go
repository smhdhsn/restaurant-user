package service

import (
	"errors"
	"strconv"

	"github.com/smhdhsn/bookstore-user/internal/model"
	"github.com/smhdhsn/bookstore-user/internal/repository/contract"
	"github.com/smhdhsn/bookstore-user/util/encrypt"
)

// This block holds public errors of this scope.
var (
	ErrParseUint = errors.New("error on parsing userID")
)

// UserService contains repositories that will be used within this service.
type UserService struct {
	uRepo contract.UserRepository
}

// NewUserService creates a user service with it's dependencies.
func NewUserService(uRepo contract.UserRepository) *UserService {
	return &UserService{
		uRepo: uRepo,
	}
}

// StoreUserReq is responsible for holding user's information to be stored into database.
type StoreUserReq struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Status    string `json:"status"`
}

// Store is responsible for storing user data inside database.
func (s *UserService) Store(req *StoreUserReq) (*model.User, error) {
	user := model.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  encrypt.ToMD5(req.Password),
		Status:    req.Status,
	}

	return s.uRepo.Store(&user)
}

// Find is responsible for storing user data inside database.
func (s *UserService) Find(ident string) (*model.User, error) {
	userID, err := strconv.ParseUint(ident, 10, 32)
	if err != nil {
		return nil, ErrParseUint
	}

	return s.uRepo.Find(uint(userID))
}

// UpdateUserReq is responsible for holding user's data to be updated inside database.
type UpdateUserReq struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Status    string `json:"status"`
}

// Update is responsible for updating user's information inside database.
func (s *UserService) Update(req *UpdateUserReq, ident string) error {
	userID, err := strconv.ParseUint(ident, 10, 32)
	if err != nil {
		return ErrParseUint
	}

	user := model.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  encrypt.ToMD5(req.Password),
		Status:    req.Status,
	}

	return s.uRepo.Update(&user, uint(userID))
}

// Destroy is responsible for deleting a user from the database.
func (s *UserService) Destroy(ident string) error {
	userID, err := strconv.ParseUint(ident, 10, 32)
	if err != nil {
		return ErrParseUint
	}

	return s.uRepo.Destroy(uint(userID))
}
