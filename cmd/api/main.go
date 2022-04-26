package main

import (
	"log"

	"github.com/smhdhsn/bookstore-user/internal/config"
	"github.com/smhdhsn/bookstore-user/internal/db"
	"github.com/smhdhsn/bookstore-user/internal/http"
	"github.com/smhdhsn/bookstore-user/internal/http/resource"
	"github.com/smhdhsn/bookstore-user/internal/repository/mysql"

	uHandler "github.com/smhdhsn/bookstore-user/internal/http/handler/user"
	uService "github.com/smhdhsn/bookstore-user/internal/service/user"
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

	// instantiate repositories
	uRepo := mysql.NewUserRepo(dbConn)

	// instantiate services
	uSourceService := uService.NewSourceService(uRepo)
	uSearchService := uService.NewSearchService(uRepo)
	uAuthService := uService.NewAuthService(uRepo)

	// instantiate handlers
	uSourceHandler := uHandler.NewSource(uSourceService)
	uSearchHandler := uHandler.NewSearch(uSearchService)
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
