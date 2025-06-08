package main

import (
	"context"
	"log"
	"time"

	pb "github.com/jayanthsharabu/grpc/avg/proto"
)

func calculateAvg(c pb.AvgServiceClient) {
	log.Println("Avg got invoked")

	stream, err := (c).AvgCal(context.Background())

	if err != nil {
		log.Fatalf("There is an error: %v", err)
	}

	nums := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, num := range nums {
		log.Printf("Sending %v\n", num)
		stream.Send(&pb.NumsRequest{
			Num: num,
		})
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response: %v", err)
	}
	log.Printf("Avg: %v\n", res.Result)

}
