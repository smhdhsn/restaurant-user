package main

import (
	"fmt"
	"log"

	"github.com/smhdhsn/bookstore-user/internal/config"
	"github.com/smhdhsn/bookstore-user/internal/db"
	"github.com/smhdhsn/bookstore-user/internal/http"
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

	httpServer := http.New()
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(httpServer.Listen(fmt.Sprintf(":%d", conf.Server.Port)))
}
