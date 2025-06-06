package main

import (
	"context"
	"log"

	pb "github.com/jayanthsharabu/grpc/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println(("doGreet got invoked"))

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		Name: "jay",
	})

	if err != nil {
		log.Fatalf("Couldnot Greet : %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)

}
