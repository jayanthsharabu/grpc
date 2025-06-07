package main

import (
	"context"
	"io"
	"log"

	pb "github.com/jayanthsharabu/grpc/greet/proto"
)

func doGreetMany(c pb.GreetServiceClient) {
	log.Println("doGreetmany got invoked")

	req := &pb.GreetRequest{
		Name: "Jay",
	}

	stream, err := (c).GreetMany(context.Background(), req)

	if err != nil {
		log.Fatalf("There is an error: %v", err)
	}
	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream")
		}

		log.Printf("Greetmany times: %v\n", msg.Result)
	}

}
