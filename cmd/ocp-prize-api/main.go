package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	"github.com/jmoiron/sqlx"
	api "github.com/ozoncp/ocp-prize-api/internal/api"
	"github.com/ozoncp/ocp-prize-api/internal/configuration"
	"github.com/ozoncp/ocp-prize-api/internal/producer"
	desc "github.com/ozoncp/ocp-prize-api/pkg/ocp-prize-api"
)

var (
	grpcPort     = ":8082"
	dbDriverName = "sqlmock"
)

func parseConfigFile(conf *configuration.Configuration, filename string) error {
	file, _ := os.Open(filename)
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(conf)
	if err != nil {
		return err
	}
	return nil
}

func run() error {

	conf := &configuration.Configuration{}
	err := parseConfigFile(conf, "conf.json")
	if err != nil {
		log.Printf("Error parsing config: %s", err.Error())
		conf = nil
	} else {
		log.Print("Config file loaded successfully")
	}
	if conf != nil {
		grpcPort = conf.GRPCPort
		dbDriverName = conf.DBDriverName
	}

	var db *sql.DB
	sqlxDB := sqlx.NewDb(db, dbDriverName)

	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	ctx := context.Background()
	var confKey configuration.ConfigurationKey = "configuration"
	ctx = context.WithValue(ctx, confKey, conf)
	prod := producer.NewProducer(ctx, "OcpPrizeApi")
	desc.RegisterOcpPrizeApiServer(s, api.NewOcpPrizeApi(ctx, sqlxDB, prod))

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
