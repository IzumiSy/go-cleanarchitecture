.PHONY: build run clean test db/migrate db/clean

build:
	go build

run: main.go
	go run main.go -http

clean:
	rm -rf go-cleanarchitecture

test/unit:
	go test -v ./...

test/integration:
	docker run --net=go-cleanarchitecture_default -it -v "$$(pwd):/app" -w /app apiaryio/dredd dredd \
		api-description.apib http://localhost:8080 --hookfiles=./dredd_hook.js

db/migrate:
	docker run --net=go-cleanarchitecture_default --rm -v "$$(pwd)/schemas/sql:/flyway/sql" -v "$$(pwd)/config:/flyway/config" \
		flyway/flyway -configFiles=/flyway/config/flyway.conf -locations=filesystem:/flyway/sql migrate

db/clean:
	docker run --net=go-cleanarchitecture_default --rm -v "$$(pwd)/schemas/sql:/flyway/sql" -v "$$(pwd)/config:/flyway/config" \
		flyway/flyway -configFiles=/flyway/config/flyway.conf -locations=filesystem:/flyway/sql clean
