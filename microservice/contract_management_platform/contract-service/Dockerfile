FROM golang:1.24-alpine

# ติดตั้ง tzdata
RUN apk add --no-cache tzdata

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o contract-server ./cmd/server

# ตั้ง timezone ภายใน container
ENV TZ=Asia/Bangkok

EXPOSE 8082

CMD ["./contract-server"]
