package http

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/smhdhsn/bookstore-user/internal/config"
	"github.com/smhdhsn/bookstore-user/internal/http/resource"
)

// Server contains server's services.
type Server struct {
	router *gin.Engine
}

// New creates a new http server.
func New(uResource *resource.UserResource) *Server {
	r := gin.New()
	apiRouter := r.Group("/api")

	uRouter := apiRouter.Group("/users")
	uRouter.POST("/", uResource.Source.Store)
	uRouter.GET("/:userID", uResource.Source.Find)
	uRouter.PUT("/:userID", uResource.Source.Update)
	uRouter.DELETE("/:userID", uResource.Source.Destroy)

	return &Server{
		r,
	}
}

// Listen is responsible for starting the HTTP server.
func (s *Server) Listen(conf config.ServerConf) error {
	return s.router.Run(fmt.Sprintf("%s:%d", conf.Host, conf.Port))
}
