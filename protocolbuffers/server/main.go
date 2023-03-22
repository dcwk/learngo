package main

import (
	"log"
	"net"
	"context"
	"google.golang.org/grpc"
)


func (s *server) AddProduct(ctx context.Context, in *pb.Product) (pb.ProductID, error) {
	return new pb.ProductID
}

func (s *server) GetProduct(ctx context.Context, in *pb.ProductID) (pb.Product, error) {
	return new pb.Product;
}

func main() {
	lis, _ := net.Listen("tcp", port)
	s := grpc.NewServer()
	pb.RegisterProductInfoServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serve: %v", err)
	}
}
