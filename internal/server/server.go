package server

import (
	"fmt"
	"net"

	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/smhdhsn/restaurant-user/internal/config"
	"github.com/smhdhsn/restaurant-user/internal/server/resource"

	log "github.com/smhdhsn/restaurant-user/internal/logger"
	uspb "github.com/smhdhsn/restaurant-user/internal/protos/user/source"
)

// Server contains server's services.
type Server struct {
	listener net.Listener
	grpc     *grpc.Server
	conf     *config.ServerConf
}

// NewServer creates a new http server.
func NewServer(c *config.ServerConf, ur *resource.UserResource) (*Server, error) {
	// Listen to a specific host and port for incoming requests.
	l, err := net.Listen(c.Protocol, fmt.Sprintf("%s:%d", c.Host, c.Port))
	if err != nil {
		return nil, errors.Wrap(err, "failed to listen to port")
	}

	// Instantiate gRPC server.
	s := grpc.NewServer()

	// Register gRPC service handlers.
	uspb.RegisterUserSourceServiceServer(s, ur.SourceHandler)

	return &Server{
		listener: l,
		grpc:     s,
		conf:     c,
	}, nil
}

// Listen is responsible for starting the HTTP server.
func (s *Server) Listen() error {
	log.Info(fmt.Sprintf("%s server started listening on port <%d>", s.conf.Protocol, s.conf.Port))
	return s.grpc.Serve(s.listener)
}
