FROM golang:latest

WORKDIR /usr/src/app

COPY . /usr/src/app
RUN GO111MODULE=on make build

CMD ["/usr/src/app/go-cleanarchitecture", "-http"]
