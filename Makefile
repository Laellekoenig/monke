run:
	go run cmd/monke/main.go

build:
	go build -o bin/monke cmd/monke/main.go

format:
	go fmt ./...
