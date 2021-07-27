.PHONY: build run test

all: build

build:
	go build -v ./cmd/server

clean:
	rm -f server
	go clean -i

run:
	go run ./cmd/main.go

test:
	go test ./test