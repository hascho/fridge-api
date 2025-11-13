# syntax=docker/dockerfile:1

FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /go-fridge ./cmd/server

FROM alpine:latest

WORKDIR /

COPY --from=builder /go-fridge /go-fridge

CMD ["./go-fridge"]
