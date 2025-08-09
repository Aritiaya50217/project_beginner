FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# เพิ่ม tzdata
RUN apk add --no-cache tzdata

RUN go build -o stock-service ./cmd/service/main.go


ENV TZ=Asia/Bangkok

CMD ["./stock-service"]
