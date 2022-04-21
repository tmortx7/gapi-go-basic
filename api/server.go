package api

import (
	"fmt"

	db "github.com/tmortx7/gapi-go-basic/db/sqlc"
	"github.com/tmortx7/gapi-go-basic/pb"
	"github.com/tmortx7/gapi-go-basic/token"
	"github.com/tmortx7/gapi-go-basic/util"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	pb.UnimplementedSimpleServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer creates a new gRPC server.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
