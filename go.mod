module github.com/ozoncp/ocp-prize-api

go 1.16

require (
	github.com/envoyproxy/protoc-gen-validate v0.6.1
	github.com/golang/glog v0.0.0-20210429001901-424d2337a529 // indirect
	github.com/golang/mock v1.5.0
	github.com/golang/protobuf v1.5.2
	github.com/golangci/golangci-lint v1.40.1 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/iancoleman/strcase v0.1.3 // indirect
	github.com/lyft/protoc-gen-star v0.5.3 // indirect
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.13.0
	github.com/ozoncp/ocp-prize-api/pkg/ocp-prize-api v0.0.0-00010101000000-000000000000
	github.com/rs/zerolog v1.22.0
	google.golang.org/genproto v0.0.0-20210611144927-798beca9d670
	google.golang.org/grpc v1.38.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0 // indirect
)

replace github.com/ozoncp/ocp-prize-api/pkg/ocp-prize-api => ./pkg/ocp-prize-api
