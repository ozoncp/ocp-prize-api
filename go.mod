module github.com/ozoncp/ocp-prize-api

go 1.16

replace github.com/ozoncp/ocp-prize-api/pkg/ocp-prize-api => ./pkg/ocp-prize-api

require (
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/Masterminds/squirrel v1.5.0
	github.com/golang/mock v1.6.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/jmoiron/sqlx v1.3.4
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.13.0
	github.com/ozoncp/ocp-prize-api/pkg/ocp-prize-api v0.0.0-00010101000000-000000000000
	github.com/rs/zerolog v1.23.0
	google.golang.org/genproto v0.0.0-20210614182748-5b3b54cad159 // indirect
	google.golang.org/grpc v1.38.0
)
