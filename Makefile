.PHONY: run test build clean fmt vet

run:
	go run .

test:
	go test ./...

build:
	mkdir -p bin
	go build -o bin/scanner .

clean:
	rm -rf bin

fmt:
	go fmt ./...

vet:
	go vet ./...
