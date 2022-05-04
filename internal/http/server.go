package http

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/smhdhsn/bookstore-user/internal/config"
	"github.com/smhdhsn/bookstore-user/internal/http/resource"
)

// Server contains server's services.
type Server struct {
	uResource *resource.UserResource
	router    *gin.Engine
}

// New creates a new http server.
func New(uResource *resource.UserResource) *Server {
	s := new(Server)
	s.uResource = uResource
	s.router = gin.New()

	pvGroup := s.router.Group("/_/")
	s.mapUserPV(pvGroup)

	pbGroup := s.router.Group("/")
	s.mapUserPB(pbGroup)

	return s
}

// mapUserPV is responsible for mapping user's sensitive routes.
func (s *Server) mapUserPV(r *gin.RouterGroup) {
	apiRouter := r.Group("/api")
	uRouter := apiRouter.Group("/users")

	uRouter.POST("/", s.uResource.Source.Store)
	uRouter.GET("/:userID", s.uResource.Source.Find)
	uRouter.PUT("/:userID", s.uResource.Source.Update)
	uRouter.DELETE("/:userID", s.uResource.Source.Destroy)

	uRouter.GET("/search", s.uResource.Search.List)
}

// mapUserPB is responsible for mapping user's public routes.
func (s *Server) mapUserPB(r *gin.RouterGroup) {
	apiRouter := r.Group("/api")
	uRouter := apiRouter.Group("/users")

	uRouter.GET("/:userCode", s.uResource.Source.Find)

	uRouter.GET("/search", s.uResource.Search.Index)
}

// Listen is responsible for starting the HTTP server.
func (s *Server) Listen(conf config.ServerConf) error {
	return s.router.Run(fmt.Sprintf("%s:%d", conf.Host, conf.Port))
}
