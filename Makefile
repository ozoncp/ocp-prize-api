run:
	go run cmd/ozon-prize-api/main.go

lint:
	golint ./...

test:
	go test -v ./...

.PHONY: build
build: vendor-proto .generate .build

PHONY: .generate
.generate:
		mkdir -p swagger
		mkdir -p pkg/ocp-prize-api
		protoc -I vendor.protogen \
				--go_out=pkg/ocp-prize-api --go_opt=paths=import \
				--go-grpc_out=pkg/ocp-prize-api --go-grpc_opt=paths=import \
				--grpc-gateway_out=pkg/ocp-prize-api \
				--grpc-gateway_opt=logtostderr=true \
				--grpc-gateway_opt=paths=import \
				--validate_out lang=go:pkg/ocp-prize-api \
				--swagger_out=allow_merge=true,merge_file_name=api:swagger \
				api/ocp-prize-api/ocp-prize-api.proto
		mv pkg/ocp-prize-api/gihtub.com/ozoncp/ocp-prize-api/pkg/ocp-prize-api/* pkg/ocp-prize-api/
		rm -rf pkg/ocp-prize-api/gihtub.com
		mkdir -p cmd/ocp-prize-api

PHONY: .build
.build:
		go build -o cmd/ozon-prize-api cmd/ozon-prize-api/main.go

PHONY: install
install: build .install

PHONY: .install
install:
		go install cmd/grpc-server/main.go

PHONY: vendor-proto
vendor-proto: .vendor-proto

PHONY: .vendor-proto
.vendor-proto:
		mkdir -p vendor.protogen
		mkdir -p vendor.protogen/api/ocp-prize-api
		cp api/ocp-prize-api/ocp-prize-api.proto vendor.protogen/api/ocp-prize-api
		@if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi


.PHONY: deps
deps: install-go-deps

.PHONY: install-go-deps
install-go-deps: .install-go-deps

.PHONY: .install-go-deps
.install-go-deps:
		ls go.mod || go mod init
		go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
		go get -u github.com/golang/protobuf/proto
		go get -u github.com/golang/protobuf/protoc-gen-go
		go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
		go install google.golang.org/grpc/cmd/protoc-gen-go-grpc