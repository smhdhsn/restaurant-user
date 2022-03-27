package http

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/smhdhsn/bookstore-user/internal/config"
	"github.com/smhdhsn/bookstore-user/internal/http/handler"
	"github.com/smhdhsn/bookstore-user/internal/service"
)

// Server contains server's services.
type Server struct {
	router *gin.Engine
}

// New creates a new http server.
func New(uServ *service.UserService) *Server {
	r := gin.New()
	apiRouter := r.Group("/api")

	uHandler := handler.NewUserHandler(uServ)

	uRouter := apiRouter.Group("/users")

	uRouter.POST("/", uHandler.Store)
	uRouter.GET("/:userID", uHandler.Find)
	uRouter.PUT("/:userID", uHandler.Update)
	uRouter.DELETE("/:userID", uHandler.Destroy)

	return &Server{
		r,
	}
}

// Listen is responsible for starting the HTTP server.
func (s *Server) Listen(conf config.ServerConf) error {
	return s.router.Run(fmt.Sprintf("%s:%d", conf.Host, conf.Port))
}
