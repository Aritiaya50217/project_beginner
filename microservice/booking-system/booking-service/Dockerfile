FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# เพิ่ม tzdata
RUN apk add --no-cache tzdata

RUN go build -o booking-service ./cmd/service/main.go

FROM alpine:latest

WORKDIR /app

# เพิ่ม tzdata ที่ production image 
RUN apk add --no-cache tzdata

COPY --from=builder /app/booking-service .

ENV TZ=Asia/Bangkok

CMD ["./booking-service"]
