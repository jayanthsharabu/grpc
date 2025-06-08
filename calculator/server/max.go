package main

import (
	"log"

	"io"

	pb "github.com/jayanthsharabu/grpc/calculator/proto"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {

	log.Printf("Max function is invoked")

	var maxi int32 = 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.Send(&pb.MaxResponse{
				Result: maxi,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		if req.Num > maxi {
			maxi = req.Num
		}
	}

	return nil

}
