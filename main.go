package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"net"
	"net/http"

	"github.com/titusdishon/simple_bank/api"
	db "github.com/titusdishon/simple_bank/db/sqlc"
	"github.com/titusdishon/simple_bank/gapi"
	"github.com/titusdishon/simple_bank/pb"
	"github.com/titusdishon/simple_bank/util"
)

func main() {
	config, err := util.LoadConfig(".", "app")
	if err != nil {
		log.Fatal(" Cannot load config: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal(" Cannot connect to db: ", err)
	}
	store := db.NewStore(conn)
	go runGatewayServer(config, store)
	runGRPCServer(config, store)

}

func runGatewayServer(config util.Config, store db.Store) {
	server, err := gapi.NewSever(config, store)
	if err != nil {
		log.Fatal(" Cannot create grpc server: ", err)
	}
	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})
	grpcMux := runtime.NewServeMux(jsonOption)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err = pb.RegisterSimpleBankHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal(" Cannot register handler server: ", err)
	}
	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	listener, err := net.Listen("tcp", config.HTTPServerAddress)
	if err != nil {
		log.Fatal(" Cannot create listener: ", err)
	}
	fmt.Printf("start HTTP server at: %s\n", listener.Addr().String())
	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatal(" Cannot start server: ", err)
	}
}

func runGRPCServer(config util.Config, store db.Store) {
	server, err := gapi.NewSever(config, store)
	if err != nil {
		log.Fatal(" Cannot create grpc server: ", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterSimpleBankServer(grpcServer, server)
	reflection.Register(grpcServer)
	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal(" Cannot create listener: ", err)
	}
	fmt.Printf("start gRPC server at: %s\n", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal(" Cannot start server: ", err)
	}
}

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewSever(config, store)
	if err != nil {
		log.Fatal(" Cannot create server: ", err)
	}
	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal(" Cannot start server: ", err)
	}
}
