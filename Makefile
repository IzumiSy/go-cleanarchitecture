.PHONY: build clean test reset migrate run

build:
	go build

run: main.go
	go run main.go -http

migrate:
	docker run --rm -v schemas/sql:/flyway/sql -v config:/flyway/config flyway/flyway -configFiles=/flyway/config/flyway.conf -locations=filesystem:/flyway/sql migrate

clean:
	rm -rf go-cleanarchitecture

test:
	go test -v ./...
