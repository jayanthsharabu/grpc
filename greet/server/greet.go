package main

import (
	"log"

	pb "github.com/jayanthsharabu/grpc/greet/proto"
	"golang.org/x/net/context"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {

	log.Printf("Greet function was invoked with", in)
	return &pb.GreetResponse{
		Result: "Hello " + in.Name,
	}, nil
}
