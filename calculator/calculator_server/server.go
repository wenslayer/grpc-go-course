package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"net"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	"github.com/wenslayer/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	log.Printf("Sum() invoked with: %v\n", req)

	operands := req.GetOperands().GetInts()

	var sum int64 = 0
	for _, num := range operands {
		sum += num
	}
	res := calculatorpb.SumResponse{
		Result: sum,
	}

	return &res, nil
}

func (*server) PrimeFactorization(req *calculatorpb.PrimeFactorizationRequest, stream calculatorpb.CalculatorService_PrimeFactorizationServer) error {
	log.Printf("PrimeFactorization() invoked with: %v\n", req)

	numberToFactor := req.GetNumber()
	var divisor int64 = 2

	// TODO: implement https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
	for numberToFactor > 1 {
		if numberToFactor%divisor == 0 {
			res := &calculatorpb.PrimeFactorizationResponse{
				Factor: divisor,
			}
			stream.Send(res)

			numberToFactor = numberToFactor / divisor
		} else {
			divisor++
		}
	}

	return nil
}

func (*server) Average(stream calculatorpb.CalculatorService_AverageServer) error {
	log.Println("Average() invoked")

	sum, count := int64(0), int64(0)
	res := &calculatorpb.AverageResponse{Average: 0}

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			if count > 0 {
				res.Average = float64(sum) / float64(count)
			}
			return stream.SendAndClose(res)
		}
		if err != nil {
			log.Fatalf("error while reading client stream: %v\n", err)
		}

		count++
		sum += req.GetNumber()
		log.Printf("%d: Sum now %d\n", count, sum)
	}
}

func (*server) Maximum(stream calculatorpb.CalculatorService_MaximumServer) error {
	log.Println("Maximum() invoked")

	res := &calculatorpb.MaximumResponse{Maximum: 0}

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}

		num := req.GetNumber()
		if num > res.Maximum {
			res.Maximum = num
			log.Printf("Max now %20d\n", res.Maximum)
			sendErr := stream.Send(res)
			if sendErr != nil {
				log.Fatalf("Error while sending data to client: %v", sendErr)
				return sendErr
			}
		}
	}
}

func (*server) SquareRoot(ctx context.Context, req *calculatorpb.SquareRootRequest) (*calculatorpb.SquareRootResponse, error) {
	log.Printf("SquareRoot() invoked with: %v\n", req)

	num := req.GetNumber()

	if num < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received negative number: %v", num),
		)
	}

	return &calculatorpb.SquareRootResponse{
		Result: math.Sqrt(float64(num)),
	}, nil
}

func main() {
	log.Println("Hello world, I'm a calculator server")

	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})
	reflection.Register(s)

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
