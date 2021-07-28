.PHONY: build clean run test

all: build

build:
	go build -v ./cmd/server

clean:
	rm -f server
	go clean -i

run:
	go run ./cmd/server/main.go

test:
	go test ./test