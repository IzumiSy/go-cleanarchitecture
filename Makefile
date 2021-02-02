.PHONY: build clean test reset migrate run

build:
	go build

run: main.go
	go run main.go

migrate:
	go run main.go --migrate

reset:
	rm -f go-cleanarchitecture.db
	make migrate

clean:
	rm -rf go-cleanarchitecture

test:
	go test -v ./...
