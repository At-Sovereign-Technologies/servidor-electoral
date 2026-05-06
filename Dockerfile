# Build stage
FROM golang:1.26 AS builder

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o bin/ ./...

# Runtime stage
FROM debian:bookworm-slim

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/bin/ .

# Expose application port
EXPOSE 8080

# Run the application
CMD ["./electoral-server"]