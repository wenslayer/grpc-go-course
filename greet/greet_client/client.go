package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/wenslayer/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello world, I'm a client.")

	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer connection.Close()

	c := greetpb.NewGreetServiceClient(connection)

	// doUnary(c)

	doServerStreaming(c)
}

// func doUnary(connection *grpc.ClientConn) {
func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Start unary RPC...")

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Marc",
			LastName:  "Wensauer",
		},
	}

	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet() RPC: %v", err)
	}

	fmt.Printf("Greet() response:\n<<<\n%v\n>>>\n", res.Result)

}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Start stream RPC...")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Marc",
			LastName:  "Wensauer",
		},
	}

	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling GreetManyTimes() RPC: %v", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// end of stream
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		log.Printf("Response from GreetManyTimes:\n>>> %v <<<\n\n", msg.GetResult())
	}
}
