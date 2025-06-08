package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/jayanthsharabu/grpc/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateBlog(ctx context.Context, In *pb.Blog) (*pb.BlogId, error) {

	log.Printf("Create Blog was invoked with %v\n", In)

	data := BlogItem{
		AuthorId: In.Authorid,
		Title:    In.Title,
		Content:  In.Content,
	}

	res, err := collection.InsertOne(ctx, data)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal Error %v\n", err),
		)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot Convert to OID"),
		)
	}

	return &pb.BlogId{
		Id: oid.Hex(),
	}, nil

}
