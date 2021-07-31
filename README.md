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
  -migrate string
    	migration mode (default "up")
```

## Run with Docker
```sh
$ docker-compose up --build
$ docker-compose run app make migrate
```

## Architecture
WIP
