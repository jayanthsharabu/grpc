package main

import (
	"log"

	pb "github.com/jayanthsharabu/grpc/primes/proto"
)

func (s *Server) Primes(in *pb.NumRequest, stream pb.PrimesService_PrimesServer) error {

	log.Printf("Primes was invoked %v\n", in)

	temp := in.Num

	k := int32(2)

	for temp > 1 {
		if temp%k == 0 {
			log.Printf("Sending %v\n", k)
			stream.Send(&pb.PrimesResponse{
				Result: k,
			})
			temp = temp / k
		} else {
			k += 1
		}

	}

	return nil

}
