FROM golang:1.24 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o contract-service ./cmd/server

FROM alpine:3.18
RUN apk add --no-cache ca-certificates
WORKDIR /root/
COPY --from=builder /app/contract-service .
ENV TZ=Asia/Bangkok
EXPOSE 8082
CMD ["./contract-service"]
