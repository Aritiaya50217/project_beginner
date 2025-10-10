# Stage 1: Build
FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# ใช้ build tags จาก docker-compose
ARG BUILD_TAGS
RUN go build -o user-service ./cmd/user-service

# Stage 2: Runtime
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/user-service .

ENV PORT=8081
EXPOSE ${PORT}

CMD ["./user-service"]
