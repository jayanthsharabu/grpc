package main

import (
	"context"
	"log"

	pb "github.com/jayanthsharabu/grpc/blog/proto"
)

func creatBlog(c pb.BlogServiceClient) string {

	log.Println("Create Blog was invoked")

	blog := &pb.Blog{
		Authorid: "Jay",
		Title:    "Sample Title",
		Content:  "Sample Content",
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Error while creating blog: %v\n", err)
	}

	log.Printf("Blog has been created : %s\n", res.Id)

	return res.Id
}
