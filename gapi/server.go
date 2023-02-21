package gapi

import (
	"fmt"

	db "github.com/titusdishon/simple_bank/db/sqlc"
	"github.com/titusdishon/simple_bank/pb"
	"github.com/titusdishon/simple_bank/token"
	"github.com/titusdishon/simple_bank/util"
)

// Server serves HTTP requests for the banking services.
type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewSever  creates a new GRPC server.
func NewSever(config util.Config, store db.Store) (*Server, error) {
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
