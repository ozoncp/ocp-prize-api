package main

import (
	"database/sql"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/jmoiron/sqlx"
	api "github.com/ozoncp/ocp-prize-api/internal/api"
	desc "github.com/ozoncp/ocp-prize-api/pkg/ocp-prize-api"
)

const (
	grpcPort = ":82"
)

func run() error {
	var db *sql.DB
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	desc.RegisterOcpPrizeApiServer(s, api.NewOcpPrizeApi(sqlxDB))

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
