all: main

.PHONY: main clean test bench install uninstall

main:
	goimports -w .
	go build

uninstall:
	go clean -i

clean:
	go clean

test:
	go test ./...

bench:
	go test -bench=. -benchmem