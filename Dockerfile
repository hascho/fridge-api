# syntax=docker/dockerfile:1

############################
# 1. Builder Stage
############################
FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /go-fridge ./cmd/server

############################
# 2. Debug Stage (for local dev only)
############################
FROM golang:1.25 AS debug

WORKDIR /app
RUN go install github.com/go-delve/delve/cmd/dlv@latest

COPY . .
RUN go mod tidy

# Build the debug binary with optimisations disabled
RUN CGO_ENABLED=0 GOOS=linux go build -gcflags="all=-N -l" -o /go-fridge ./cmd/server

# Delve listens on port 40000 by default
EXPOSE 40000
CMD ["dlv", "exec", "/go-fridge", "--headless", "--listen=:40000", "--api-version=2", "--accept-multiclient"]

############################
# 3. Production Stage
############################
FROM alpine:latest AS prod

WORKDIR /

COPY --from=builder /go-fridge /go-fridge

CMD ["./go-fridge"]
