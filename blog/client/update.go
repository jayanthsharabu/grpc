package main

import (
	"context"
	"log"

	pb "github.com/jayanthsharabu/grpc/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {

	log.Println("Update Blog was invoked")

	new_blog := &pb.Blog{
		Id:       id,
		Authorid: "Not Jay",
		Title:    "New title",
		Content:  "New content",
	}

	_, err := c.UpdateBlog(context.Background(), new_blog)

	if err != nil {
		log.Fatalf("Error while Updating blog: %v\n", err)
	}

	log.Printf("Blog has been Updated")

}
