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
- Uses [earthly](https://github.com/earthly/earthly) for repeatable build

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
$ earthly +middlewares-up
$ earthly +db-migrate
$ earthly +run
```

## Tests
```sh
# Unit testing
$ earthly +unit-test

# Integration testing
$ earthly +integration-test
```

## Architecture
WIP
