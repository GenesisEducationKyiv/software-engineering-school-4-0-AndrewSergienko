FROM golang:latest as builder

WORKDIR /app

COPY . .

RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN go mod download

RUN go build -o /goapp ./cmd
