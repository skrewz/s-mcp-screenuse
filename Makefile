.PHONY: build run test clean

build: s-mcp-screenuse

s-mcp-screenuse: main.go go.mod go.sum
	go build -o s-mcp-screenuse .

run: build
	./s-mcp-screenuse

test:
	go test -v ./...

clean:
	rm -f s-mcp-screenuse
	go clean
