FROM golang:1.17-alpine3.15 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/api/main.go

CMD ["./.bin/app"]