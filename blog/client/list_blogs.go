package main

import (
	"context"
	"io"
	"log"

	pb "github.com/jayanthsharabu/grpc/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func listBlogs(c pb.BlogServiceClient) {

	log.Println("List Blogs was invoked")

	stream, err := c.ListBlog(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalf("error while streaming %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something happened %v\n", err)
		}

		log.Println(res)
	}

}
