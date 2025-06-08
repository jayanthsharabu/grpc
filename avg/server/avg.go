package main

import (
	"io"
	"log"

	pb "github.com/jayanthsharabu/grpc/avg/proto"
)

func (s *Server) AvgCal(stream pb.AvgService_AvgCalServer) error {

	log.Printf("AvgCal was invoked")

	sum := int32(0)
	count := int32(0)

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Result: sum / count,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		sum += req.Num
		count++
	}

}
