package main

import (
	"log"
	"net"

	pb "github.com/jayanthsharabu/grpc/primes/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.PrimesServiceServer
}

var addr string = "0.0.0.0:50051"

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on %v\n", addr)
	}

	log.Printf("Listening on %v\n", addr)

	s := grpc.NewServer()
	pb.RegisterPrimesServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v\n", addr)
	}
}
