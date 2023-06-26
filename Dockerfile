FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go. ./

expose 8080

CMD["go","run","main.go"]