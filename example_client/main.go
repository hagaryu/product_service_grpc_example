package main

import (
	"context"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "grpc_example/product_service_proto/v1"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewProductServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Send multiple products
	products := []*pb.Product{
		{Name: "Product 1", Price: 100.0},
		{Name: "Product 2", Price: 200.0},
		{Name: "Product 3", Price: 300.0},
	}

	for _, product := range products {
		r, err := c.AddProduct(ctx, product)
		if err != nil {
			log.Fatalf("could not save product: %v", err)
		}
		log.Printf("Server response: %s", r.GetMessage())
	}

	// Get all products
	r, err := c.GetAllProducts(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("could not get products: %v", err)
	}
	log.Printf("Products: %v", r.GetProducts())
}
