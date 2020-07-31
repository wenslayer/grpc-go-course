package main

import (
	"fmt"
	"log"

	"github.com/wenslayer/grpc-go-course/greet/greetpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello, I'm a client.")

	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer connection.Close()

	c := greetpb.NewGreetServiceClient(connection)

	fmt.Printf("Created client: %f", c)
}
