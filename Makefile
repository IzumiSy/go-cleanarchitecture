.PHONY: build clean test reset migrate run

build:
	go build

run: main.go
	go run main.go -http

clean:
	rm -rf go-cleanarchitecture

test:
	go test -v ./...
