package api

import (
	"github.com/tmortx7/gapi-go-basic/pb"
	"github.com/tmortx7/gapi-go-basic/util"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	pb.UnimplementedSimpleServer
	config util.Config
}
