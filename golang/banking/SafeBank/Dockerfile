# Dockerfile
FROM golang:1.23-alpine

WORKDIR /app

# Copy go.mod + go.sum first and download deps
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build app
RUN go build -o main ./cmd/main.go

EXPOSE 8080

# Run the app
CMD ["./main"]
