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

run:
  BUILD +image
  COPY docker-compose.yml .
  WITH DOCKER --compose docker-compose.yml --load app:latest=+image
    docker run --net=go-cleanarchitecture-network --env APP_ENV=production --rm app:latest -http
  END

test:
  BUILD +unit-test
  BUILD +integration-test

unit-test:
  FROM +deps
  COPY . .
  RUN go test -v ./...

integration-test:
  BUILD +image
  COPY docker-compose.yml .
  COPY dredd_hook.js api-description.apib .
  WITH DOCKER \
      --compose docker-compose.yml \
      --load app:latest=+image \
      --pull apiaryio/dredd \
      --pull flyway/flyway:7
    RUN sleep 15 && \
      docker run --net=go-cleanarchitecture-network --rm -v "$(pwd)/schemas/sql:/flyway/sql" -v "$(pwd)/config:/flyway/conf" flyway/flyway:7 migrate && \
      docker run --net=go-cleanarchitecture-network --env APP_ENV=production --rm app:latest -http & \
      docker run --net=go-cleanarchitecture-network --rm -it -v "$(pwd):/app" -w /app apiaryio/dredd dredd \
		    api-description.apib localhost:8080 --hookfiles=./dredd_hook.js
  END
