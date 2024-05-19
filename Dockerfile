FROM golang:latest

ENV GOPATH /go

ENV PATH $GOPATH/bin:$PATH

WORKDIR $GOPATH/src/app

COPY . .

RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN go mod download
RUN go build -o service-app