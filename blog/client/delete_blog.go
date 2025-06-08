package main

import (
	"context"
	"log"

	pb "github.com/jayanthsharabu/grpc/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func deleteBlog(c pb.BlogServiceClient, id string) (*emptypb.Empty, error) {

	log.Println("delete was invoked")

	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})

	if err != nil {
		log.Fatalf("error while deleting the object %v\n", err)
	}

	log.Println("It was deleted")

	return &emptypb.Empty{}, nil

}
