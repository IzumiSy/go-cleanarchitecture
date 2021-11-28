VERSION 0.6
FROM golang:1.17-alpine3.14
WORKDIR /go-cleanarchitecture

deps:
  COPY go.mod go.sum .
  RUN apk add --no-cache build-base
  RUN go mod download

build:
  FROM +deps
  COPY . .
  RUN go build -o build/go-cleanarchitecture main.go
  SAVE ARTIFACT build/go-cleanarchitecture AS LOCAL build/go-cleanarchitecture

image:
  COPY +build/go-cleanarchitecture .
  EXPOSE 8080
  ENTRYPOINT ["/go-cleanarchitecture/go-cleanarchitecture"]
  SAVE IMAGE go-cleanarchitecture:latest

test:
  FROM +unit-test
  FROM +integration-test

unit-test:
  FROM +deps
  COPY . .
  RUN go test -v ./...

integration-test:
  FROM +image
  COPY . .
  WITH DOCKER \
      --compose docker-compose.yml \
      --load app:latest=+image \
      # --pull earthly/dind:latest \ # for caching
      --pull flyway/flyway:7       # for caching
    RUN sleep 15 && \
      docker run --net=go-cleanarchitecture-network --rm -v "$(pwd)/schemas/sql:/flyway/sql" -v "$(pwd)/config:/flyway/config" \
		    flyway/flyway:7 -configFiles=/flyway/config/flyway.conf -locations=filesystem:/flyway/sql migrate
        # docker run --net=go-cleanarchitecture-network --env APP_ENV=production --rm app:latest -http
  END

all:
  FROM +deps
  FROM +build
  FROM +docker
