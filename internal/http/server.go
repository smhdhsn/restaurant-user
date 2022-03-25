package http

import (
	"github.com/gin-gonic/gin"
	"github.com/smhdhsn/bookstore-user/internal/http/handler"
)

// Server contains server's services.
type Server struct {
	router *gin.Engine
}

// New creates a new http server.
func New() *Server {
	r := gin.New()

	uHandler := handler.NewUserHandler()

	r.POST("/users", uHandler.Create)
	r.GET("/users/:userID", uHandler.Get)
	r.PUT("/users/:userID", uHandler.Update)
	r.PATCH("/users/:userID", uHandler.Update)
	r.DELETE("/users/:userID", uHandler.Delete)
	r.GET("/internal/users/search", uHandler.Search)
	r.POST("/users/login", uHandler.Login)

	return &Server{
		r,
	}
}

// Listen is responsible for starting the HTTP server.
func (s *Server) Listen(port string) error {
	return s.router.Run(port)
}
