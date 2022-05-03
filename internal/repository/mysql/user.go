package mysql

import (
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/smhdhsn/bookstore-user/internal/model"
	"github.com/smhdhsn/bookstore-user/internal/repository/contract"
	"gorm.io/gorm"
)

// UserRepo contains repository's database connection.
type UserRepo struct {
	model *model.User
	db    *gorm.DB
}

// NewUserRepo creates an instance of the repository with database connection.
func NewUserRepo(db *gorm.DB, m *model.User) contract.UserRepository {
	return &UserRepo{
		model: m,
		db:    db,
	}
}

// Store is responsible for storing a user inside database.
func (r *UserRepo) Store(u *model.UserDTO) (*model.UserDTO, error) {
	err := r.db.Model(r.model).Create(u).Error
	if err != nil {
		if err.(*mysql.MySQLError).Number == contract.DuplicateEntryErrNum {
			return nil, contract.ErrDuplicateEntry
		}

		return nil, err
	}

	return u, nil
}

// FindBy is responsible for finding a record with a given column matching a given value.
func (r *UserRepo) FindBy(req contract.FilterBy) (model.UserDTOList, error) {
	uList := make(model.UserDTOList, 0)
	err := r.db.Model(r.model).Where(fmt.Sprintf("%s = ?", req.Field), req.Value).Find(&uList).Error
	if err != nil {
		return nil, err
	} else if len(uList) == 0 {
		return nil, contract.ErrRecordNotFound
	}

	return uList, nil
}

// Inspect is responsible for fetching user's full details from database.
func (r *UserRepo) Inspect(userID uint) (*model.UserDTO, error) {
	u := new(model.UserDTO)
	err := r.db.Model(r.model).First(u, userID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, contract.ErrRecordNotFound
		}

		return nil, err
	}

	return u, nil
}

// Find is responsible for fetching user's limited details from database.
func (r *UserRepo) Find(userID uint) (*model.UserDTO, error) {
	u := new(model.UserDTO)
	err := r.db.Model(r.model).Select("id", "status", "created_at", "updated_at").First(u, userID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, contract.ErrRecordNotFound
		}

		return nil, err
	}

	return u, nil
}

// Update is responsible for updating user's information inside database.
func (r *UserRepo) Update(user *model.UserDTO, userID uint) error {
	tx := r.db.Model(r.model).Where("id = ?", userID).Updates(user)
	if tx.Error != nil && tx.Error.(*mysql.MySQLError).Number == contract.DuplicateEntryErrNum {
		return contract.ErrDuplicateEntry
	} else if tx.RowsAffected == 0 {
		return contract.ErrRecordNotFound
	}

	return tx.Error
}

// Destroy is responsible for deleting a user from the database.
func (r *UserRepo) Destroy(userID uint) error {
	tx := r.db.Where("id = ?", userID).Delete(r.model)
	if err := tx.Error; err != nil {
		return err
	} else if tx.RowsAffected == 0 {
		return contract.ErrRecordNotFound
	}

	return nil
}
