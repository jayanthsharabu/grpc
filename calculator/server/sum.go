package main

import (
	"log"

	pb "github.com/jayanthsharabu/grpc/Calculator/proto"
	"golang.org/x/net/context"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {

	log.Printf("Calculator function is invoked with %v\n", in)
	return &pb.SumResponse{

		Result: in.Num1 + in.Num2,
	}, nil

}
