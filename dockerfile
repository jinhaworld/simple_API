FROM golang:1.15

WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -o serverd ./cmd/...