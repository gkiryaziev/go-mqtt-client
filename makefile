all: main

.PHONY: main install uninstall clean test bench lint

main:
	go build

uninstall:
	go clean -i

clean:
	go clean

test:
	go test ./...

bench:
	go test -bench=. -benchmem

lint:
	go fmt ./...
	goimports -w .
	golint ./...
	go vet ./...
