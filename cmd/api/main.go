package main

import (
	"log"

	"github.com/smhdhsn/bookstore-user/internal/config"
	"github.com/smhdhsn/bookstore-user/internal/db"
	"github.com/smhdhsn/bookstore-user/internal/http"
	"github.com/smhdhsn/bookstore-user/internal/repository/mysql"
	"github.com/smhdhsn/bookstore-user/internal/service"
)

// main is the application's kernel.
func main() {
	// read configurations
	conf, err := config.LoadConf()
	if err != nil {
		log.Fatal(err)
	}

	// create a database connection
	dbConn, err := db.Connect(conf.DB)
	if err != nil {
		log.Fatal(err)
	}

	// initialize auto migration
	if err := db.InitMigrations(dbConn); err != nil {
		log.Fatal(err)
	}

	// instanciate repositories
	uRepo := mysql.NewUserRepo(dbConn)

	// instanciate services
	uServ := service.NewUserService(uRepo)

	// instanciate handlers
	httpServer := http.New(uServ)
	if err != nil {
		log.Fatal(err)
	}

	// listen and serve
	log.Fatal(httpServer.Listen(conf.Server))
}
