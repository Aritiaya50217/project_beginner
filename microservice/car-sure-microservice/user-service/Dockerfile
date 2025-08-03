FROM golang:1.23-alpine

WORKDIR /app

# คัดลอก go.mod และ go.sum จาก root
COPY ./go.mod ./go.sum ./
RUN go mod download

# คัดลอกเฉพาะ user-service และ internal ที่จำเป็น
COPY ./user-service ./user-service
COPY ./config ./config

# เข้าไปที่ user-service
WORKDIR /app/user-service

# Build ไฟล์ main.go
RUN go build -o main ./cmd/main.go

# Run binary
CMD ["./main"]
