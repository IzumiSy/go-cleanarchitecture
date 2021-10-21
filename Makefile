.PHONY: build run clean test db/migrate db/clean

build:
	go build

run: main.go
	go run main.go -http

clean:
	rm -rf go-cleanarchitecture

test:
	go test -v ./...

db/migrate:
	docker run --net=go-cleanarchitecture_default --rm \ 
		-v $$(pwd)/schemas/sql:/flyway/sql -v $$(pwd)/config:/flyway/config flyway/flyway \ 
		-configFiles=/flyway/config/flyway.conf -locations=filesystem:/flyway/sql migrate

db/clean:
	docker run --net=go-cleanarchitecture_default --rm \ 
		-v $$(pwd)/schemas/sql:/flyway/sql -v $$(pwd)/config:/flyway/config flyway/flyway \ 
		-configFiles=/flyway/config/flyway.conf -locations=filesystem:/flyway/sql clean
