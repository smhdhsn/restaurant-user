package mysql

import (
	"gorm.io/gorm"
)

// migrationModels holds the schema of the models to be migrated to database.
var migrationModels = [...]interface{}{
	&user{},
}

// InitMigrations migrates models to the database.
func InitMigrations(db *gorm.DB) error {
	return db.AutoMigrate(migrationModels[:]...)
}
