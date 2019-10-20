.PHONY: reset migrate run

run: main.go
	go run main.go

migrate:
	go run main.go --migrate

reset:
	rm -f go-cleanarchitecture.db
	make migrate


