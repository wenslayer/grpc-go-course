package main

import (
	"context"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc/codes"

	"github.com/wenslayer/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	log.Printf("Greet() invoked with: %v\n", req)

	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()

	result := "Guten tag, Herr Doktor Diplomingenieur " + lastName + ". May I call you " + firstName + "?"
	res := greetpb.GreetResponse{
		Result: result,
	}

	return &res, nil
}

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes() invoked with: %v\n", req)
	firstName := req.GetGreeting().GetFirstName()

	for i := 0; i < 10; i++ {
		result := "Hello, " + firstName + " number " + strconv.Itoa(i)
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}

	return nil
}

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	log.Printf("LongGreet() invoked\n")

	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// finished reading from client
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		firstName := req.GetGreeting().GetFirstName()
		result += "Guten tag, " + firstName + "!\n"
	}
}

func (*server) GreetEveryone(stream greetpb.GreetService_GreetEveryoneServer) error {
	log.Printf("GreetEveryone() invoked\n")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}

		firstName := req.GetGreeting().GetFirstName()
		result := "Guten tag, " + firstName + "!\n"
		sendErr := stream.Send(&greetpb.GreetEveryoneResponse{
			Result: result,
		})
		if sendErr != nil {
			log.Fatalf("Error while sending data to client: %v", sendErr)
			return sendErr
		}
	}
}

func (*server) GreetWithDeadline(ctx context.Context, req *greetpb.GreetWithDeadlineRequest) (*greetpb.GreetWithDeadlineResponse, error) {
	log.Printf("GreetWithDeadline() invoked with: %v\n", req)

	numContextChecks := 3
	for i := 1; i <= numContextChecks; i++ {
		log.Printf("...check context [%d/%d]...\n", i, numContextChecks)
		if ctx.Err() != nil {
			log.Printf("...context error detected: >>> %v <<<\n", ctx.Err())
			return nil, status.Error(codes.Canceled, "client canceled the request")
		}
		time.Sleep(1 * time.Second)
	}
	log.Println("...context checked")

	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()

	result := "Guten tag, Herr Doktor Diplomingenieur " + lastName + ". May I call you " + firstName + "?"
	res := &greetpb.GreetWithDeadlineResponse{
		Result: result,
	}

	return res, nil
}

// HostAndPort set at build time
var HostAndPort = "localhost:12345"

// SSLEnabled set at build time
var SSLEnabled = "false"

// ServerCertFile set at build time
var ServerCertFile = ""

// ServerKeyFile set at build time
var ServerKeyFile = ""

func main() {
	log.Println("Hello world, I'm a greet server")

	log.Printf("...listen on [%v]...\n", HostAndPort)
	listener, err := net.Listen("tcp", HostAndPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}

	useTLS, err := strconv.ParseBool(SSLEnabled)
	if err != nil {
		log.Fatalf("Could not convert to boolean: %v", SSLEnabled)
	}
	if useTLS {
		log.Println("...secure communication ENABLED")
		creds, sslErr := credentials.NewServerTLSFromFile(ServerCertFile, ServerKeyFile)
		if sslErr != nil {
			log.Fatalf("Failed to load certificates: %v", sslErr)
			return
		}
		opts = append(opts, grpc.Creds(creds))
	} else {
		log.Println("...secure communication DISABLED")
	}

	s := grpc.NewServer(opts...)
	greetpb.RegisterGreetServiceServer(s, &server{})
	log.Println("...server registered; ready for connections...")
	reflection.Register(s)
	log.Println("...reflection registered...")

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	log.Println("...goodbye")
}
