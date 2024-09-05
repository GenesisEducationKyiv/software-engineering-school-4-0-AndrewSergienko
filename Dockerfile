FROM golang:latest as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /app/notifier ./cmd/notifier
RUN go build -o /app/rateservice ./cmd/rateservice

RUN go install github.com/pressly/goose/v3/cmd/goose@v3.21.1

RUN apt-get update && \
    apt-get install -y curl && \
    curl -fsSL https://just.systems/install.sh | bash -s -- --to /usr/local/bin/

CMD sh -c "just ./scripts/${APP_NAME}/migrate & ./\${APP_NAME}"
