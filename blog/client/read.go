package main

import (
	"context"
	"log"

	pb "github.com/jayanthsharabu/grpc/blog/proto"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {

	log.Println("Read Blog was invoked")

	blog := &pb.BlogId{
		Id: id,
	}

	res, err := c.ReadBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Error while Reading blog: %v\n", err)
	}

	log.Printf("Blog has been retured : %v\n", res)

	return res
}
