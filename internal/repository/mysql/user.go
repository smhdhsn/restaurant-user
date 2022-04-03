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
	db *gorm.DB
}

// NewUserRepo creates an instance of the repository with database connection.
func NewUserRepo(db *gorm.DB) contract.UserRepository {
	return &UserRepo{db}
}

// Find is responsible for finding user with given ID inside database.
func (r *UserRepo) Find(userID uint) (*model.User, error) {
	user := new(model.User)
	err := r.db.First(user, userID).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, contract.ErrRecordNotFound
	}
	return user, err
}

// FindBy is responsible for finding a record with a given column matching a given value.
func (r *UserRepo) FindBy(req contract.FilterBy) ([]*model.User, error) {
	users := make([]*model.User, 0)
	err := r.db.Where(fmt.Sprintf("%s = ?", req.Field), req.Value).Find(users).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, contract.ErrRecordNotFound
	}
	return users, err
}

// Store is responsible for storing a user inside database.
func (r *UserRepo) Store(user *model.User) error {
	err := r.db.Create(user).Error
	if err != nil && err.(*mysql.MySQLError).Number == contract.DuplicateEntryErrNum {
		return contract.ErrDuplicateEntry
	}
	return err
}

// Update is responsible for updating user's information inside database.
func (r *UserRepo) Update(user *model.User, userID uint) error {
	tx := r.db.Where("id = ?", userID).Updates(user)
	if tx.RowsAffected == 0 {
		return contract.ErrRecordNotFound
	} else if tx.Error != nil && tx.Error.(*mysql.MySQLError).Number == contract.DuplicateEntryErrNum {
		return contract.ErrDuplicateEntry
	}
	return tx.Error
}

// Destroy is responsible for deleting a user from the database.
func (r *UserRepo) Destroy(userID uint) error {
	tx := r.db.Where("id = ?", userID).Delete(new(model.User))
	if tx.RowsAffected == 0 {
		return contract.ErrRecordNotFound
	}
	return tx.Error
}
