FROM golang:latest

WORKDIR /usr/src/app

COPY . /usr/src/app
RUN GO111MODULE=on make build

ENTRYPOINT ["/usr/src/app/go-cleanarchitecture", "-http"]
