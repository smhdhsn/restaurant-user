package main

import (
	"log"

	"github.com/smhdhsn/bookstore-user/internal/config"
	"github.com/smhdhsn/bookstore-user/internal/db"
	"github.com/smhdhsn/bookstore-user/internal/http"
	"github.com/smhdhsn/bookstore-user/internal/repository/mysql"
	"github.com/smhdhsn/bookstore-user/internal/service"
)

func main() {
	conf, err := config.LoadConf()
	if err != nil {
		log.Fatal(err)
	}

	dbConn, err := db.Connect(conf.DB)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.InitMigrations(dbConn); err != nil {
		log.Fatal(err)
	}

	uRepo := mysql.NewUserRepo(dbConn)

	uServ := service.NewUserService(uRepo)

	httpServer := http.New(uServ)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(httpServer.Listen(conf.Server))
}
