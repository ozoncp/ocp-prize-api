package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	api "github.com/ozoncp/ocp-prize-api/internal/api"
	"github.com/ozoncp/ocp-prize-api/internal/configuration"
	"github.com/ozoncp/ocp-prize-api/internal/producer"
	desc "github.com/ozoncp/ocp-prize-api/pkg/ocp-prize-api"
)

var (
	configFilename = "config.json.template"
	grpcPort       = ":8082"
	dbDriverName   = "pgx"
	dbHost         = "postgres"
	dbPort         = 5432
	dbLogin        = ""
	dbPassword     = ""
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
	err := parseConfigFile(conf, configFilename)
	if err != nil {
		log.Printf("Error parsing config: %s", err.Error())
		conf = nil
		return err
	} else {
		log.Print("Config file loaded successfully")
	}
	if conf != nil {
		grpcPort = conf.GRPCPort
		dbDriverName = conf.DBDriverName
		dbHost = conf.DBHost
		dbPort = conf.DBPort
		dbLogin = conf.DBLogin
		dbPassword = conf.DBPassword
	}

	dbInfo := fmt.Sprintf("postgres://%s:%s@%s:%d/postgres?sslmode=disable",
		dbLogin, dbPassword, dbHost, dbPort)
	sqlxDB, err := sqlx.Connect(dbDriverName, dbInfo)
	if err != nil {
		log.Fatalf("failed to init database: %v", err)
	} else {
		log.Print("Database connected successfully")
	}

	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Printf("Start listening port: %s", grpcPort)
	}

	s := grpc.NewServer()
	ctx := context.Background()
	var confKey configuration.ConfigurationKey = "configuration"
	ctx = context.WithValue(ctx, confKey, conf)
	prod := producer.NewProducer(ctx, "OcpPrizeApi")
	desc.RegisterOcpPrizeApiServer(s, api.NewOcpPrizeApi(ctx, sqlxDB, prod))

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	} else {
		log.Print("GRPC server started successfully")
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
