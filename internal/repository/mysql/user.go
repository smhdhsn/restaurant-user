package mysql

import (
	"errors"

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
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, contract.ErrRecordNotFound
	}
	return user, nil
}

// Store is responsible for storing a user inside database.
func (r *UserRepo) Store(user *model.User) (err error) {
	err = r.db.Create(user).Error
	if err != nil && err.(*mysql.MySQLError).Number == contract.DuplicateEntryErrNum {
		return contract.ErrDuplicateEntry
	}
	return
}

// Update is responsible for updating user's information inside database.
func (r *UserRepo) Update(user *model.User, userID uint) error {
	tx := r.db.Where("id = ?", userID).Updates(user)
	if tx.RowsAffected == 0 {
		return contract.ErrRecordNotFound
	} else if err := tx.Error; err.(*mysql.MySQLError).Number == contract.DuplicateEntryErrNum {
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
