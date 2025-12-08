FROM golang:1.24-alpine

# ติดตั้ง tzdata
RUN apk add --no-cache tzdata

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o notification-worker ./cmd/service

# ตั้ง timezone ภายใน container
ENV TZ=Asia/Bangkok

CMD ["./notification-worker"]
