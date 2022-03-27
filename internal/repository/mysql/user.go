package mysql

import (
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

// Store is responsible for storing a user inside database.
func (r *UserRepo) Store(user *model.User) (*model.User, error) {
	tx := r.db.Model(new(model.User)).Create(user)
	return user, tx.Error
}

// Update is responsible for updating user's information inside database.
func (r *UserRepo) Update(user *model.User, userID uint) error {
	return r.db.Where("id = ?", userID).Updates(user).Error
}

// Find is responsible for finding user with given ID inside database.
func (r *UserRepo) Find(userID uint) (*model.User, error) {
	user := new(model.User)
	tx := r.db.Where("id = ?", userID).First(user)
	return user, tx.Error
}

// Destroy is responsible for deleting a user from the database.
func (r *UserRepo) Destroy(userID uint) error {
	return r.db.Where("id = ?", userID).Delete(new(model.User)).Error
}
