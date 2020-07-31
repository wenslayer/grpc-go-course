package main

import (
	"fmt"
	"log"
	"net"

	"github.com/wenslayer/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	fmt.Println("hello world")

	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
