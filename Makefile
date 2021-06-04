build:
	go build -o cmd/ozon-prize-api cmd/ozon-prize-api/main.go

run:
	go run cmd/ozon-prize-api/main.go

lint:
	golint ./...

test:
	go test -v ./...