package api

import (
	db "github.com/tmortx7/gapi-go-basic/db/sqlc"
	"github.com/tmortx7/gapi-go-basic/pb"
	"github.com/tmortx7/gapi-go-basic/util"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	pb.UnimplementedSimpleServer
	config util.Config
	store  db.Store
}

// NewServer creates a new gRPC server.
func NewServer(config util.Config, store db.Store) (*Server, error) {

	server := &Server{
		config: config,
		store:  store,
	}

	return server, nil
}
