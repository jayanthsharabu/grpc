package main

import (
	"context"
	"io"
	"log"

	pb "github.com/jayanthsharabu/grpc/primes/proto"
)

func calculatePrimes(c pb.PrimesServiceClient) {
	log.Println("Primes got invoked")

	req := &pb.NumRequest{
		Num: 100,
	}

	stream, err := (c).Primes(context.Background(), req)

	if err != nil {
		log.Fatalf("There is an error: %v", err)
	}
	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream")
		}

		log.Printf("Primes: %v\n", msg.Result)
	}

}
