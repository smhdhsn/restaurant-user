package handler

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/smhdhsn/restaurant-user/internal/repository/entity"
	"github.com/smhdhsn/restaurant-user/internal/service/dto"

	authProto "github.com/smhdhsn/restaurant-user/internal/protos/user/auth"
	serviceContract "github.com/smhdhsn/restaurant-user/internal/service/contract"
)

// AuthHandler contains services that can be used within user auth handler.
type AuthHandler struct {
	authServ serviceContract.UserAuthService
}

// NewAuthHandler creates a new user auth handler.
func NewAuthHandler(as serviceContract.UserAuthService) authProto.UserAuthServiceServer {
	return &AuthHandler{
		authServ: as,
	}
}

// Store is responsible for storing a user into database.
func (s *AuthHandler) Store(ctx context.Context, req *authProto.UserStoreRequest) (*authProto.UserStoreResponse, error) {
	uDTO := singleStoreReqToUserDTO(req)

	uEntity, err := s.authServ.Store(uDTO)
	if err != nil {
		if errors.Is(err, serviceContract.ErrDuplicateEntry) {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}

		return nil, status.Errorf(codes.Internal, "internal server error: %w", err)
	}

	resp := singleUserEntityToStoreResp(uEntity)

	return resp, nil
}

// singleStoreReqToUserDTO is responsible for transforming a Store request into user dto struct.
func singleStoreReqToUserDTO(req *authProto.UserStoreRequest) *dto.User {
	return &dto.User{
		FirstName: req.GetFirstName(),
		LastName:  req.GetLastName(),
		Email:     req.GetEmail(),
		Password:  req.GetPassword(),
	}
}

// singleUserEntityToStoreResp is responsible for transforming a user entity into Store response struct.
func singleUserEntityToStoreResp(uEntity *entity.User) *authProto.UserStoreResponse {
	return &authProto.UserStoreResponse{
		Id:        uEntity.ID,
		FirstName: uEntity.FirstName,
		LastName:  uEntity.LastName,
		Email:     uEntity.Email,
	}
}

// FindBy is responsible for fetching user's details from database.
func (s *AuthHandler) FindBy(ctx context.Context, req *authProto.UserFindByRequest) (*authProto.UserFindByResponse, error) {
	uDTO := singleFindByReqToUserDTO(req)

	uEntity, err := s.authServ.FindBy(uDTO)
	if err != nil {
		if errors.Is(err, serviceContract.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		return nil, status.Errorf(codes.Internal, "internal server error: %w", err)
	}

	resp := singleUserEntityToFindByResp(uEntity)

	return resp, nil
}

// singleFindByReqToUserDTO is responsible for transforming a FindBy request into user dto struct.
func singleFindByReqToUserDTO(req *authProto.UserFindByRequest) *dto.User {
	return &dto.User{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}
}

// singleUserEntityToFindByResp is responsible for transforming a user entity into FindBy response struct.
func singleUserEntityToFindByResp(uEntity *entity.User) *authProto.UserFindByResponse {
	return &authProto.UserFindByResponse{
		Id:        uEntity.ID,
		FirstName: uEntity.FirstName,
		LastName:  uEntity.LastName,
		Email:     uEntity.Email,
	}
}
