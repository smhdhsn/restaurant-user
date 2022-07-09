package mysql

import (
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/smhdhsn/restaurant-user/internal/repository/entity"

	repositoryContract "github.com/smhdhsn/restaurant-user/internal/repository/contract"
)

// user represents the users table on database.
type user struct {
	ID        uint32 `gorm:"primarykey"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Email     string `gorm:"size:255;uniqueIndex;not null"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// UserRepo contains repository's database connection.
type UserRepo struct {
	model user
	db    *gorm.DB
}

// NewUserRepository creates an instance of the repository with database connection.
func NewUserRepository(db *gorm.DB) repositoryContract.UserRepository {
	return &UserRepo{
		model: user{},
		db:    db,
	}
}

// FindBy is responsible for fetching user's details from database.
func (r *UserRepo) FindBy(uEntity *entity.User) (*entity.User, error) {
	uModel := singleUserEntityToModel(uEntity)

	err := r.db.Model(r.model).
		Where("email = ?", uModel.Email).
		Where("password = ?", uModel.Password).
		Take(uModel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repositoryContract.ErrRecordNotFound
		}

		return nil, errors.Wrap(err, "error on fetching user by email and password from database")
	}

	uEntity = singleUserModelToEntity(uModel)

	return uEntity, nil
}

// Store is responsible for storing a user inside database.
func (r *UserRepo) Store(uEntity *entity.User) (*entity.User, error) {
	uModel := singleUserEntityToModel(uEntity)

	err := r.db.Model(r.model).Create(uModel).Error
	if err != nil {
		if err.(*mysql.MySQLError).Number == repositoryContract.DuplicateEntryErrNum {
			return nil, repositoryContract.ErrDuplicateEntry
		}

		return nil, errors.Wrap(err, "error on storing user's information inside database")
	}

	uEntity = singleUserModelToEntity(uModel)

	return uEntity, nil
}

// singleUserEntityToModel is responsible for transforming a user entity to user model struct.
func singleUserEntityToModel(uEntity *entity.User) *user {
	return &user{
		ID:        uEntity.ID,
		FirstName: uEntity.FirstName,
		LastName:  uEntity.LastName,
		Email:     uEntity.Email,
		Password:  uEntity.Password,
		CreatedAt: uEntity.CreatedAt,
		UpdatedAt: uEntity.UpdatedAt,
	}
}

// singleUserModelToEntity is responsible for transforming a user entity to user model struct.
func singleUserModelToEntity(uModel *user) *entity.User {
	return &entity.User{
		ID:        uModel.ID,
		FirstName: uModel.FirstName,
		LastName:  uModel.LastName,
		Email:     uModel.Email,
		Password:  uModel.Password,
		CreatedAt: uModel.CreatedAt,
		UpdatedAt: uModel.UpdatedAt,
	}
}
