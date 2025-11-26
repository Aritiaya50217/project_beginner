FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o contract-proxy ./cmd/server

EXPOSE 8083

CMD ["./contract-proxy"]
