package main

import (
	"context"
	"log"
	"net"

	"github.com/wenslayer/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	log.Printf("Sum() invoked with: %v\n", req)

	operands := req.GetOperands().GetInts()

	var sum int32 = 0
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

	prime := req.GetNumber()
	var factor int32 = 2

	for prime > 1 {
		if prime%factor == 0 {
			res := &calculatorpb.PrimeFactorizationResponse{
				Factor: factor,
			}
			stream.Send(res)

			prime = prime / factor
		} else {
			factor++
		}
	}

	return nil
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
