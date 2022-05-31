package mysql

import (
	"errors"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"

	"github.com/smhdhsn/restaurant-user/internal/model"

	repositoryContract "github.com/smhdhsn/restaurant-user/internal/repository/contract"
)

// UserRepo contains repository's database connection.
type UserRepo struct {
	model model.User
	db    *gorm.DB
}

// NewUserRepository creates an instance of the repository with database connection.
func NewUserRepository(db *gorm.DB, m model.User) repositoryContract.UserRepository {
	return &UserRepo{
		model: m,
		db:    db,
	}
}

// Store is responsible for storing a user inside database.
func (r *UserRepo) Store(u *model.UserDTO) (*model.UserDTO, error) {
	err := r.db.Model(r.model).Create(u).Error
	if err != nil {
		if err.(*mysql.MySQLError).Number == repositoryContract.DuplicateEntryErrNum {
			return nil, repositoryContract.ErrDuplicateEntry
		}

		return nil, err
	}

	return u, nil
}

// Find is responsible for fetching user's full details from database.
func (r *UserRepo) Find(u *model.UserDTO) (*model.UserDTO, error) {
	err := r.db.Model(r.model).First(u, u.ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repositoryContract.ErrRecordNotFound
		}

		return nil, err
	}

	return u, nil
}

// Destroy is responsible for deleting a user from the database.
func (r *UserRepo) Destroy(u *model.UserDTO) error {
	tx := r.db.Where("id = ?", u.ID).Delete(r.model)
	if err := tx.Error; err != nil {
		return err
	} else if tx.RowsAffected == 0 {
		return repositoryContract.ErrRecordNotFound
	}

	return nil
}

// Update is responsible for updating user's information inside database.
func (r *UserRepo) Update(u *model.UserDTO) error {
	tx := r.db.Model(r.model).Where("id = ?", u.ID).Updates(u)
	if err := tx.Error; err != nil {
		if err.(*mysql.MySQLError).Number == repositoryContract.DuplicateEntryErrNum {
			return repositoryContract.ErrDuplicateEntry
		} else {
			return err
		}
	} else if tx.RowsAffected == 0 {
		return repositoryContract.ErrRecordNotFound
	}

	return nil
}
