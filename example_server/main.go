package main

import (
	"context"
	"log"
	"net"

	pb "grpc_example/product_service_proto/v1"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedProductServiceServer
	products []*pb.Product
}

func (s *server) AddProduct(ctx context.Context, in *pb.Product) (*pb.ProductResponse, error) {
	product := &pb.Product{Name: in.GetName(), Price: in.GetPrice()}
	s.products = append(s.products, product)
	return &pb.ProductResponse{Message: "Product saved successfully"}, nil
}

func (s *server) GetAllProducts(ctx context.Context, in *pb.Empty) (*pb.ProductsResponse, error) {
	return &pb.ProductsResponse{Products: s.products}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProductServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
