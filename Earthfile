VERSION 0.6
FROM golang:1.17-alpine3.14
WORKDIR /go-cleanarchitecture

RUN apk add --no-cache build-base

deps:
  COPY go.mod go.sum ./
  RUN go mod download

build:
  FROM +deps
  COPY . .
  RUN GO111MODULE=on go build -o build/go-cleanarchitecture main.go
  SAVE ARTIFACT build/go-cleanarchitecture /go-cleanarchitecture AS LOCAL build/go-cleanarchitecture

docker:
  FROM +build
  COPY +build/go-cleanarchitecture .
  EXPOSE 8080
  ENTRYPOINT ["/go-cleanarchitecture/go-cleanarchitecture"]
  SAVE IMAGE go-cleanarchitecture:latest

all:
  FROM +deps
  FROM +build
  FROM +docker
