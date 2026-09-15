.PHONY: run test build clean fmt vet check

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

check: fmt vet test
