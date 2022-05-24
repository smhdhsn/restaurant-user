package main

import (
	"github.com/smhdhsn/restaurant-user/internal/config"
	"github.com/smhdhsn/restaurant-user/internal/db"
	"github.com/smhdhsn/restaurant-user/internal/http"
	"github.com/smhdhsn/restaurant-user/internal/http/resource"
	"github.com/smhdhsn/restaurant-user/internal/model"
	"github.com/smhdhsn/restaurant-user/internal/repository/mysql"

	uHand "github.com/smhdhsn/restaurant-user/internal/http/handler/user"
	log "github.com/smhdhsn/restaurant-user/internal/logger"
	uServ "github.com/smhdhsn/restaurant-user/internal/service/user"
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
	uSourceService := uServ.NewSourceService(uRepo)

	// instantiate handlers.
	uSourceHandler := uHand.NewSourceHandler(uSourceService)

	// instantiate resources.
	uRes := resource.NewUserResource(uSourceHandler)

	// instantiate gRPC server.
	s, err := http.NewServer(&conf.Server, uRes)
	if err != nil {
		log.Fatal(err)
	}

	// listen and serve.
	if err := s.Listen(); err != nil {
		log.Fatal(err)
	}
}
