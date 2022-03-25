package db

import (
	"fmt"

	"github.com/smhdhsn/bookstore-user/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connect creates a database connection.
func Connect(c config.DBConf) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		c.User,
		c.Pass,
		c.Host,
		c.Port,
		c.Name,
	)

	return gorm.Open(mysql.Open(dsn))
}
