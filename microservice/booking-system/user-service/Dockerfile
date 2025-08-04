FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o user-service ./cmd/service/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/user-service .

CMD ["./user-service"]
