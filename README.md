# go-cleanarchitecture

[![CircleCI](https://circleci.com/gh/IzumiSy/go-cleanarchitecture/tree/master.svg?style=svg)](https://circleci.com/gh/IzumiSy/go-cleanarchitecture/tree/master)

This exploration project includes:

- Authorization flow
- Transaction
- Logging
- Asynchronous Pub/Sub (Redis, goroutine)
- Type-safety (No `interface{}` as much as possible)
- Multiple drivers (migration, web, CLI)
- Testing (unit-testing, integration-testing with dredd)
- CI integration (CircleCI)

## Build
```sh
$ make build
```

## Run
```
$ ./go-cleanarchitecture -help
Usage of ./go-cleanarchitecture:
  -http
    	http server mode
  -migrate
    	migration mode
```

## Run with Docker
```sh
# Launches all (including app)
$ docker-compose --profile app up --build
$ make db/migrate

# Launches only middlewares (for development)
$ docker-compose up --profile middleware up --build -d
$ make db/migrate
$ make run # Runs app locally (not on Docker)
```

## Tests
```sh
$ make test/unit
$ make test/integration # requires app running on Docker with docker-compose
```

## Architecture
WIP
