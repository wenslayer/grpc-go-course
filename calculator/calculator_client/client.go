package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/wenslayer/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello world, I'm a client.")

	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer connection.Close()

	c := calculatorpb.NewCalculatorServiceClient(connection)

	// doSum(c)
	doPrimeFactorization(c)
}

func doSum(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Start sum RPC...")

	req := &calculatorpb.SumRequest{
		Operands: &calculatorpb.Operands{
			Ints: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Sum RPC: %v", err)
	}

	fmt.Printf("Sum() response: <<< %v >>>\n", res.Result)
}

func doPrimeFactorization(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Start PrimeFactorization stream RPC...")

	// var number int64 = int64Pow(2, 16) - 15 // 65521
	// var number int64 = 104729
	// var number int64 = 15485863
	// var number int64 = 512927357
	// var number int64 = int64Pow(2, 32) - 5 // 4294967291
	var number int64 = int64Pow(2, 33) - 9 // 8589934583
	// var number int64 = 7898765131657

	req := &calculatorpb.PrimeFactorizationRequest{
		Number: number,
	}

	resStream, err := c.PrimeFactorization(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling PrimeFactorization() RPC: %v", err)
	}

	var factors []int64

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// end of stream
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		factor := msg.GetFactor()
		fmt.Printf("Response from PrimeFactorization: >>> %v <<<\n", factor)
		factors = append(factors, factor)
	}

	fmt.Printf("Factors of %d: %v\n", number, factors)
}

func int64Pow(x int, y int) int64 {
	var result int64 = 1
	for i := 1; i <= y; i++ {
		result = result * int64(x)
	}
	return result
}
