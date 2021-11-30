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
  LOCALLY
  WITH DOCKER --load app:latest=+image
    RUN docker run --net=go-cleanarchitecture-network --net-alias=app \
      -p 8080:8080 --env APP_ENV=production --rm app:latest -http
  END

db-migrate:
  LOCALLY
  WITH DOCKER --pull flyway/flyway:7
    RUN docker run --net=go-cleanarchitecture-network \
      -v "$(pwd)/schemas/sql:/flyway/sql" -v "$(pwd)/config:/flyway/conf" --rm flyway/flyway:7 migrate
  END

db-clean:
  LOCALLY
  WITH DOCKER --pull flyway/flyway:7
    RUN docker run --net=go-cleanarchitecture-network \
      -v "$(pwd)/schemas/sql:/flyway/sql" -v "$(pwd)/config:/flyway/config" --rm flyway/flyway:7 clean
  END

test:
  BUILD +unit-test
  BUILD +integration-test

unit-test:
  FROM +deps
  COPY . .
  RUN go test -v ./...

# Requires middlwares up with docker-compose
integration-test:
  LOCALLY
  BUILD +db-migrate
  WITH DOCKER --load app:latest=+image --pull apiaryio/dredd
    RUN cid=`docker run -d --net=go-cleanarchitecture-network --net-alias=app --env APP_ENV=production --rm app:latest -http` && \
      docker run --net=go-cleanarchitecture-network -v "$(pwd):/app" -w /app --rm apiaryio/dredd dredd \
        api-description.apib http://app:8080 --hookfiles=./dredd_hook.js && \
      docker stop $cid
  END
