# Multi-stage Dockerfile for building and running the servidor-electoral Go service

FROM golang:1.21 AS builder
WORKDIR /src

# Download dependencies first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the full project then build the server binary
COPY . .

# Build static binary for Linux
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags "-s -w" -o /app/server ./cmd/electoral-server

FROM alpine:3.18
RUN apk add --no-cache ca-certificates

WORKDIR /app

# Copy binary and required assets (templates/static) from builder
COPY --from=builder /app/server ./server
COPY --from=builder /src/templates ./templates
COPY --from=builder /src/static ./static

EXPOSE 8080

ENV PORT=8080

ENTRYPOINT ["/app/server"]
