# =========================
# Build stage
# =========================
FROM golang:1.23-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum, download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy all source code
COPY . .

# Build binary
RUN go build -o user-service ./cmd/user-service

# =========================
# Final stage
# =========================
FROM alpine:3.18

# Install tzdata, certificates, mysql client (optional)
RUN apk add --no-cache tzdata ca-certificates mysql-client bash

# Set working directory
WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/user-service .

# Expose port
EXPOSE 8081

# Run the binary directly
CMD ["./user-service"]
