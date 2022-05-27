package handler

import (
	"context"
	"errors"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/smhdhsn/restaurant-user/internal/model"
	"github.com/smhdhsn/restaurant-user/internal/repository/mysql"
	"github.com/smhdhsn/restaurant-user/internal/service"

	uspb "github.com/smhdhsn/restaurant-user/internal/protos/user/source"
)

// SourceHandler contains services that can be used within user source handler.
type SourceHandler struct {
	sourceServ service.UserSourceService
}

// NewSourceHandler creates a new user source handler.
func NewSourceHandler(sourceServ service.UserSourceService) *SourceHandler {
	return &SourceHandler{
		sourceServ: sourceServ,
	}
}

// Store is responsible for storing a user into database.
func (s *SourceHandler) Store(ctx context.Context, req *uspb.UserStoreRequest) (*uspb.UserStoreResponse, error) {
	uReq := &model.UserDTO{
		FirstName: req.GetFirstName(),
		LastName:  req.GetLastName(),
		Email:     req.GetEmail(),
		Password:  req.GetPassword(),
		Status:    req.GetStatus().String(),
	}

	uDTO, err := s.sourceServ.Store(uReq)
	if err != nil {
		if errors.Is(err, mysql.ErrDuplicateEntry) {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}

		return nil, status.Errorf(codes.Internal, "internal server error: %w", err)
	}

	resp := &uspb.UserStoreResponse{
		Id:        uDTO.ID,
		FirstName: uDTO.FirstName,
		LastName:  uDTO.LastName,
		Email:     uDTO.Email,
		Status:    uspb.Status(uspb.Status_value[strings.ToUpper(uDTO.Status)]),
		CreatedAt: uDTO.CreatedAt.Unix(),
		UpdatedAt: uDTO.UpdatedAt.Unix(),
	}

	return resp, nil
}

// Find is responsible for fetching user's details from database.
func (s *SourceHandler) Find(ctx context.Context, req *uspb.UserFindRequest) (*uspb.UserFindResponse, error) {
	uReq := &model.UserDTO{
		ID: req.GetId(),
	}

	uDTO, err := s.sourceServ.Find(uReq)
	if err != nil {
		if errors.Is(err, mysql.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		return nil, status.Errorf(codes.Internal, "internal server error: %w", err)
	}

	resp := &uspb.UserFindResponse{
		Id:        uDTO.ID,
		FirstName: uDTO.FirstName,
		LastName:  uDTO.LastName,
		Email:     uDTO.Email,
		Status:    uspb.Status(uspb.Status_value[strings.ToUpper(uDTO.Status)]),
		CreatedAt: uDTO.CreatedAt.Unix(),
		UpdatedAt: uDTO.UpdatedAt.Unix(),
	}

	return resp, nil
}

// Destroy is responsible for deleting a user from database.
func (s *SourceHandler) Destroy(ctx context.Context, req *uspb.UserDestroyRequest) (*uspb.UserDestroyResponse, error) {
	uReq := &model.UserDTO{
		ID: req.GetId(),
	}

	err := s.sourceServ.Destroy(uReq)
	if err != nil {
		if errors.Is(err, mysql.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		return nil, status.Errorf(codes.Internal, "internal server error: %w", err)
	}

	resp := &uspb.UserDestroyResponse{
		Status: true,
	}

	return resp, nil
}

// Update is responsible for updating a user's information inside database.
func (s *SourceHandler) Update(ctx context.Context, req *uspb.UserUpdateRequest) (*uspb.UserUpdateResponse, error) {
	uReq := &model.UserDTO{
		ID:        req.GetId(),
		FirstName: req.GetFirstName(),
		LastName:  req.GetLastName(),
		Email:     req.GetEmail(),
		Password:  req.GetPassword(),
		Status:    req.GetStatus().String(),
	}

	err := s.sourceServ.Update(uReq)
	if err != nil {
		if errors.Is(err, mysql.ErrDuplicateEntry) {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		} else if errors.Is(err, mysql.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		return nil, status.Errorf(codes.Internal, "internal server error: %w", err)
	}

	resp := &uspb.UserUpdateResponse{
		Status: true,
	}

	return resp, nil
}
