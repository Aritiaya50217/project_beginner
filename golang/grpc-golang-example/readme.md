ติดตั้ง protoc และ plugin สำหรับ Go
    
    brew install protobuf

    # Ubuntu/Debian
    sudo apt install -y protobuf-compiler

    # ติดตั้ง Go plugin
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

    # เพิ่ม GOPATH/bin เข้า PATH
    export PATH="$PATH:$(go env GOPATH)/bin"


Compile .proto เป็น Go

    protoc --go_out=. --go-grpc_out=. proto/helloworld.proto

วิธีการรัน
    
    1. รัน server:
        go run server/main.go

    2. เปิดอีก terminal แล้วรัน client:
        go run client/main.go
