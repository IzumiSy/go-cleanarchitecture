FROM golang:latest

ENV APP_ENV=production

WORKDIR /usr/src/app

COPY . /usr/src/app
RUN GO111MODULE=on make build

EXPOSE 8080

CMD ["/usr/src/app/go-cleanarchitecture", "-http"]
