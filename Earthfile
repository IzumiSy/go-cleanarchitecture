VERSION 0.6
FROM golang:1.17-alpine3.14
WORKDIR /go-cleanarchitecture

deps:
  COPY go.mod go.sum .
  RUN apk add --no-cache build-base
  RUN go mod download
  SAVE IMAGE --cache-hint

build:
  FROM +deps
  COPY . .
  RUN go build -o build/go-cleanarchitecture main.go
  SAVE ARTIFACT build/go-cleanarchitecture AS LOCAL build/go-cleanarchitecture

image:
  COPY +build/go-cleanarchitecture .
  EXPOSE 8080
  ENTRYPOINT ["/go-cleanarchitecture/go-cleanarchitecture"]
  SAVE IMAGE --push ghcr.io/izumisy/go-cleanarchitecture:cache

images:
  BUILD +image
  BUILD +db
  BUILD +pubsub
  BUILD +migrater
  BUILD +dredd


# Development

run:
  LOCALLY
  WITH DOCKER --load app:latest=+image
    RUN docker run --net=go-cleanarchitecture-network --net-alias=app \
      -p 8080:8080 --env APP_ENV=production --rm app:latest -http
  END

middlewares-up:
  LOCALLY
  WITH DOCKER \
      --load db:latest=+db \
      --load redis:latest=+pubsub
    RUN docker network create go-cleanarchitecture-network && \
      docker run -d --net=go-cleanarchitecture-network --net-alias=db \
        --name go-cleanarchictecture-db -p 3306:3306 --rm db:latest && \
      docker run -d --net=go-cleanarchitecture-network --net-alias=redis \
        --name go-cleanarchictecture-redis -p 6379:6379 --rm redis:latest
  END

middlewares-down:
  LOCALLY
  WITH DOCKER
    RUN docker stop go-cleanarchictecture-db go-cleanarchictecture-redis && \
      docker network rm go-cleanarchitecture-network
  END

db-migrate:
  LOCALLY
  WITH DOCKER --load migrater:latest=+migrater
    RUN docker run --net=go-cleanarchitecture-network --rm migrater:latest migrate
  END

db-clean:
  LOCALLY
  WITH DOCKER --load migrater:latest=+migrater
    RUN docker run --net=go-cleanarchitecture-network --rm migrater:latest clean
  END

# Tests

test:
  BUILD +unit-test
  BUILD +integration-test

unit-test:
  FROM +build
  RUN go test ./...

integration-test:
  LOCALLY
  WITH DOCKER \
      --load db:latest=+db \
      --load redis:latest=+pubsub \
      --load migrater:latest=+migrater \
      --load dredd:latest=+dredd \
      --load app:latest=+image
    RUN docker network create test-network && \
      docker run -d --name=db --net=test-network --net-alias=db -p 3306:3306 --rm db:latest && \
      docker run -d --name=redis --net=test-network --net-alias=redis -p 6379:6379 --rm redis:latest && \
      while ! nc 127.0.0.1 3306; do sleep 1 && echo "wait..."; done && sleep 15 && \
      docker run --net=test-network --rm migrater:latest migrate && \
      docker run -d --name=app --net=test-network --net-alias=app -p 8080:8080 --env APP_ENV=production --rm app:latest -http && \
      while ! nc 127.0.0.1 8080; do sleep 1 && echo "wait..."; done && sleep 15 && \
      docker run --net=test-network -w /app --rm dredd:latest && \
      docker stop db redis app && \
      docker network rm test-network
  END

dredd:
  FROM apiaryio/dredd
  COPY . /app
  COPY api-description.apib dredd_hook.js .
  ENTRYPOINT dredd api-description.apib http://app:8080 --hookfiles=dredd_hook.js
  SAVE IMAGE --push ghcr.io/izumisy/go-cleanarchitecture-dredd:cache

# Middlewares

db:
  FROM mysql:5.7
  ENV MYSQL_ROOT_USER=root
  ENV MYSQL_ROOT_PASSWORD=password
  ENV MYSQL_DATABASE=todoapp
  EXPOSE 3306
  SAVE IMAGE --push ghcr.io/izumisy/go-cleanarchitecture-db:cache

pubsub:
  FROM redis:6.2.6-alpine3.15
  EXPOSE 6379
  SAVE IMAGE --push ghcr.io/izumisy/go-cleanarchitecture-pubsub:cache

migrater:
  FROM flyway/flyway:7
  COPY ./config /flyway/conf
  COPY ./schemas /flyway/sql
  SAVE IMAGE --push ghcr.io/izumisy/go-cleanarchitecture-migrater:cache
