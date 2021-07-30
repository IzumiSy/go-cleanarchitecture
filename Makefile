.PHONY: build clean test reset migrate run

build:
	go build

run: main.go
	go run main.go

migrate:
	go run main.go --migrate=up

reset:
	rm -f go-cleanarchitecture.db
	sudo rm -rf tmp/

clean:
	rm -rf go-cleanarchitecture

test:
	go test -v ./...
