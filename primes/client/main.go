package main

import (
	"log"

	pb "github.com/jayanthsharabu/grpc/primes/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials((insecure.NewCredentials())))

	if err != nil {
		log.Fatalf("Failed to dail at %v\n", addr)
	}

	defer conn.Close()

	c := pb.NewPrimesServiceClient(conn)

	calculatePrimes(c)

}
