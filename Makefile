.PHONY: db/migrate db/clean

db/migrate:
	docker run --net=go-cleanarchitecture-network --rm -v "$$(pwd)/schemas/sql:/flyway/sql" -v "$$(pwd)/config:/flyway/config" \
		flyway/flyway -configFiles=/flyway/config/flyway.conf -locations=filesystem:/flyway/sql migrate

db/clean:
	docker run --net=go-cleanarchitecture-network --rm -v "$$(pwd)/schemas/sql:/flyway/sql" -v "$$(pwd)/config:/flyway/config" \
		flyway/flyway -configFiles=/flyway/config/flyway.conf -locations=filesystem:/flyway/sql clean
