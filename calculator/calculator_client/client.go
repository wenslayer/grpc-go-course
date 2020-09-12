package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"time"

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
	// doPrimeFactorization(c)
	// doAverage(c)
	doMaximum(c)
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

	// var number int64 = 104729
	// var number int64 = 15485863
	// var number int64 = 512927357
	// var number int64 = 7898765131657

	// From https://primes.utm.edu/lists/2small/0bit.html
	// var number int64 = int64Pow(2, 16) - 15 // 65521
	// var number int64 = int64Pow(2, 32) - 5 // 4294967291
	var number int64 = int64Pow(2, 33) - 9 // 8589934583

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

func doAverage(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Start client stream Average() RPC...")

	numbers := []int64{24125, 30240, 24724, 17204, 26260, 8592, 9956, 22152, 11852, 22334}

	req := &calculatorpb.AverageRequest{Number: 0}

	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatalf("error while calling Average(): %v\n", err)
		return
	}

	for _, num := range numbers {
		fmt.Printf("Send req: %v\n", num)
		req.Number = num
		stream.Send(req)
		time.Sleep(100 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while reading: %v", err)
	}
	fmt.Printf("average: %v\n", res.GetAverage())
}

func doMaximum(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Start client stream Maximum() RPC...")

	req := &calculatorpb.MaximumRequest{Number: 0}

	stream, err := c.Maximum(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
		return
	}

	// Establish channel for waiting
	waitc := make(chan struct{})

	go func() {
		for {
			num := rand.Int63()
			fmt.Printf("Send req: %20d\n", num)
			req.Number = num
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

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
			fmt.Printf("\nReceived: %20d\n\n", res.GetMaximum())
		}
		close(waitc)

	}()

	<-waitc
}

func int64Pow(x int, y int) int64 {
	var result int64 = 1
	for i := 1; i <= y; i++ {
		result = result * int64(x)
	}
	return result
}
