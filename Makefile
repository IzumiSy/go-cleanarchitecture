run: main.go
	go run main.go

migrate: migration.go
	go run migration.go

reset:
	rm -f go-cleanarchitecture.db
	go run migration.go
