package main

import (
	"github.com/smhdhsn/restaurant-user/internal/config"
	"github.com/smhdhsn/restaurant-user/internal/db"
	"github.com/smhdhsn/restaurant-user/internal/http"
	"github.com/smhdhsn/restaurant-user/internal/http/resource"
	"github.com/smhdhsn/restaurant-user/internal/model"
	"github.com/smhdhsn/restaurant-user/internal/repository/mysql"

	log "github.com/smhdhsn/restaurant-user/internal/logger"

	uHandler "github.com/smhdhsn/restaurant-user/internal/http/handler/user"
	uService "github.com/smhdhsn/restaurant-user/internal/service/user"
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

	// instantiate models
	uModel := new(model.User)

	// instantiate repositories
	uRepo := mysql.NewUserRepo(dbConn, uModel)

	// instantiate services
	uSourceService := uService.NewSourceService(uRepo)
	uSearchService := uService.NewSearchService(uRepo)
	uAuthService := uService.NewAuthService(uRepo)

	// instantiate handlers
	uSourceHandler := uHandler.NewSource(uSourceService, conf.Hash)
	uSearchHandler := uHandler.NewSearch(uSearchService, conf.Hash)
	uAuthHandler := uHandler.NewAuth(uAuthService)

	// instantiate resources
	uResource := resource.NewUserResource(uSourceHandler, uSearchHandler, uAuthHandler)

	// instantiate server
	httpServer := http.New(uResource)
	if err != nil {
		log.Fatal(err)
	}

	// listen and serve
	log.Fatal(httpServer.Listen(conf.Server))
}
