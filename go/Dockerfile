# syntax=docker/dockerfile:1

FROM golang:1.23.4-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /go-gorm-test

EXPOSE 8080

CMD [ "/go-gorm-test" ]

