package main

import (
	"fmt"
	"log"

	pb "github.com/jayanthsharabu/grpc/greet/proto"
)

func (s *Server) GreetMany(in *pb.GreetRequest, stream pb.GreetService_GreetManyServer) error {

	log.Printf("Greet Many times func was invoked: %v\n", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, number %d ", in.Name, i)

		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}

	return nil
}
