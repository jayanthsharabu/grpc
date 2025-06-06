package main

import (
	"context"
	"log"

	pb "github.com/jayanthsharabu/grpc/Calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println(("doSum got invoked"))

	res, err := c.Sum(context.Background(), &pb.SumRequest{
		Num1: 12,
		Num2: 13,
	})

	if err != nil {
		log.Fatalf("Couldnot Sum : %v\n", err)
	}

	log.Printf("Calculated: %s\n", res.Result)

}
