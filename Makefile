run:
	go run ./cmd/telenotify

build:
	go build -o ./bin/telenotify ./cmd/telenotify/main.go 

migrate:
	./migrate.sh

protoc:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
		pkg/api/grpc/notify.proto
