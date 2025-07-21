package main

import (
	"context"
	"log"
	"time"

	pb "grpc-golang-example/grpc/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloRequest{
		Name: "ChatGPT",
	})

	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Greeting : %s ", r.Message)
}
