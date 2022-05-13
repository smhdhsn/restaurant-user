package db

import (
	"github.com/smhdhsn/restaurant-user/internal/model"
	"gorm.io/gorm"
)

// MigrationModels holds the schema of the models to be migrated to database.
var MigrationModels = [...]interface{}{
	&model.User{},
}

// InitMigrations migrates models to the database.
func InitMigrations(db *gorm.DB) error {
	return db.AutoMigrate(MigrationModels[:]...)
}
