package model

import (
	"time"

	"github.com/smhdhsn/bookstore-user/internal/config"
	"github.com/smhdhsn/bookstore-user/util/encryption"
)

// User represents the users table on database.
type User struct {
	ID        uint   `gorm:"primaryKey"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Email     string `gorm:"size:255;uniqueIndex;not null"`
	Status    string `gorm:"type:enum('active', 'inactive');not null"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// UserDTOList holds a slice of type UserDTO struct.
type UserDTOList []*UserDTO

// UserDTO represents user's data transfer object.
type UserDTO struct {
	ID        uint
	FirstName string
	LastName  string
	Email     string
	Status    string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// userExternalResp is the response schema for external scope.
type userExternalResp struct {
	UserCode  string    `json:"user_code"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ToExternalResp is the response schema for external scope.
func (s *UserDTO) ToExternalResp(conf config.HashConf) userExternalResp {
	userCode, _ := encryption.EncodeHashIDs(s.ID, conf.Alphabet, conf.Salt, conf.MinLength)

	return userExternalResp{
		UserCode:  userCode,
		Status:    s.Status,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}
}

// ToExternalResp is the response schema for external scope.
func (s *UserDTOList) ToExternalResp(conf config.HashConf) []userExternalResp {
	uList := make([]userExternalResp, len(*s))
	for i, u := range *s {
		uList[i] = u.ToExternalResp(conf)
	}

	return uList
}

// userInternalResp is the response schema for internal scope.
type userInternalResp struct {
	ID        uint      `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ToInternalResp is the response schema for internal scope.
func (s *UserDTO) ToInternalResp() userInternalResp {
	return userInternalResp{
		ID:        s.ID,
		FirstName: s.FirstName,
		LastName:  s.LastName,
		Email:     s.Email,
		Status:    s.Status,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}
}

// ToInternalResp is the response schema for internal scope.
func (s *UserDTOList) ToInternalResp() []userInternalResp {
	uList := make([]userInternalResp, len(*s))
	for i, u := range *s {
		uList[i] = u.ToInternalResp()
	}

	return uList
}
