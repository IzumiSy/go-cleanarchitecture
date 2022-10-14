# go-cleanarchitecture
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
# Builds a binary
$ earthly +build

# Builds an image
$ earthly +image
```

## Run with middlewares
```sh
# Runs an application with middlewares up
$ earthly +middlewares-up
$ earthly +db-migrate
$ earthly +run

# Shuts down middlewares
$ earthly +middlewares-down
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
