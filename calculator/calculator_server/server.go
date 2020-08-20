package main

import (
	"context"
	"io"
	"log"
	"net"

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
				res.Average = float32(sum) / float32(count)
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

func main() {
	log.Println("Hello world, I'm a calculator server")

	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
