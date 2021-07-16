# go-cleanarchitecture

[![CircleCI](https://circleci.com/gh/IzumiSy/go-cleanarchitecture/tree/master.svg?style=svg)](https://circleci.com/gh/IzumiSy/go-cleanarchitecture/tree/master)

This exploration project includes:

- Authorization flow
- Transaction
- Logging
- Asynchronous Pub/Sub (WIP)
- Type-safety (No `interface{}` as much as possible)
- Multiple drivers (migration, web, CLI)

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

## Architecture
WIP
