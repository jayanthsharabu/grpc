package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/jayanthsharabu/grpc/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Println(("doSum got invoked"))

	stream, err := c.Max(context.Background())

	if err != nil {
		log.Fatalf("Couldnot Max : %v\n", err)
	}

	waitc := make(chan struct{})

	go func() {
		numbers := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		for _, num := range numbers {
			stream.Send(&pb.MaxRequest{
				Num: num,
			})
			time.Sleep(1 * time.Second)
			log.Printf("Sent: %v\n", num)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving response: %v", err)
			}
			log.Printf("Received: %v\n", res.Result)
		}
		close(waitc)
	}()

	<-waitc
}
