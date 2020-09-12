package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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
	// doServerStreaming(c)
	// doClientStreaming(c)
	// doBiDiStreaming(c)

	doUnaryWithDeadline(c, 5*time.Second)
	doUnaryWithDeadline(c, 1*time.Second)
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
	fmt.Println("Start server stream RPC...")

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

func doClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Start client stream RPC...")

	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Larry",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Curly",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Moe",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Mary",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Flo",
			},
		},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("error while calling LongGreet(): %v", err)
	}

	for _, req := range requests {
		fmt.Printf("Send req: %v\n", req)
		stream.Send(req)
		time.Sleep(100 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while reading: %v", err)
	}
	fmt.Printf("response:\n%v\n", res.GetResult())
}

func doBiDiStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Start BiDi stream RPC...")

	// create stream
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
		return
	}

	requests := []*greetpb.GreetEveryoneRequest{
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Larry",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Curly",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Moe",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Mary",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Flo",
			},
		},
	}

	waitc := make(chan struct{})

	// send messages
	go func() {
		for _, req := range requests {
			fmt.Printf("Send message: %v\n", req)
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	// receive messages
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving: %v", err)
				break
			}
			fmt.Printf("Received: %v\n", res.GetResult())
		}
		close(waitc)
	}()

	// block
	<-waitc
}

// func doUnaryWithDeadline(connection *grpc.ClientConn) {
func doUnaryWithDeadline(c greetpb.GreetServiceClient, timeout time.Duration) {
	fmt.Println("Start UnaryWithDeadline RPC...")

	req := &greetpb.GreetWithDeadlineRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Marc",
			LastName:  "Wensauer",
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	fmt.Printf("...context: >>> %v <<<\n", ctx)

	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {
		fmt.Println("...error detected...")
		grpcErr, ok := status.FromError(err)
		if ok {
			fmt.Printf("...response grpc error message: >>> %v <<<\n", grpcErr.Message())
			fmt.Printf("...response grpc error code:    >>> %v <<<\n", grpcErr.Code())

			if grpcErr.Code() == codes.DeadlineExceeded {
				fmt.Println("...timeout hit; deadline exceeded")
			} else {
				fmt.Printf("...unexpected error: %v\n", grpcErr)
			}
		} else {
			log.Fatalf("error while calling GreetWithDeadline() RPC: %v", err)
		}
		return
	}

	fmt.Printf("GreetWithDeadline() response: >>> %v <<<\n", res.Result)
}
