.PHONY: test/integration db/migrate db/clean

test/integration:
	docker run --net=go-cleanarchitecture-network --rm -it -v "$$(pwd):/app" -w /app apiaryio/dredd dredd \
		api-description.apib http://app:8080 --hookfiles=./dredd_hook.js

db/migrate:
	docker run --net=go-cleanarchitecture-network --rm -v "$$(pwd)/schemas/sql:/flyway/sql" -v "$$(pwd)/config:/flyway/config" \
		flyway/flyway -configFiles=/flyway/config/flyway.conf -locations=filesystem:/flyway/sql migrate

db/clean:
	docker run --net=go-cleanarchitecture-network --rm -v "$$(pwd)/schemas/sql:/flyway/sql" -v "$$(pwd)/config:/flyway/config" \
		flyway/flyway -configFiles=/flyway/config/flyway.conf -locations=filesystem:/flyway/sql clean
