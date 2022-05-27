package main

import (
	"github.com/smhdhsn/restaurant-user/internal/config"
	"github.com/smhdhsn/restaurant-user/internal/db"
	"github.com/smhdhsn/restaurant-user/internal/model"
	"github.com/smhdhsn/restaurant-user/internal/repository/mysql"
	"github.com/smhdhsn/restaurant-user/internal/server"
	"github.com/smhdhsn/restaurant-user/internal/server/handler"
	"github.com/smhdhsn/restaurant-user/internal/server/resource"
	"github.com/smhdhsn/restaurant-user/internal/service"

	log "github.com/smhdhsn/restaurant-user/internal/logger"
)

// main is the application's kernel.
func main() {
	// read configurations.
	conf, err := config.LoadConf()
	if err != nil {
		log.Fatal(err)
	}

	// create a database connection.
	dbConn, err := db.Connect(&conf.DB)
	if err != nil {
		log.Fatal(err)
	}

	// initialize auto migration.
	if err := db.InitMigrations(dbConn); err != nil {
		log.Fatal(err)
	}

	// instantiate models.
	uModel := new(model.User)

	// instantiate repositories.
	uRepo := mysql.NewUserRepository(dbConn, *uModel)

	// instantiate services.
	uSourceService := service.NewSourceService(uRepo)

	// instantiate handlers.
	uSourceHandler := handler.NewSourceHandler(uSourceService)

	// instantiate resources.
	uRes := resource.NewUserResource(uSourceHandler)

	// instantiate gRPC server.
	s, err := server.NewServer(&conf.Server, uRes)
	if err != nil {
		log.Fatal(err)
	}

	// listen and serve.
	if err := s.Listen(); err != nil {
		log.Fatal(err)
	}
}
