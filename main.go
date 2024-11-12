package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "grpc-example/products"

	"google.golang.org/grpc"
)

type Product struct {
	Ean            string
	Sku            string
	BrutPrice      int32
	NetPrice       int32
	AdditionalInfo string
	Quantity       int32
}

var fakeDB = map[string]Product{
	"PROD001": {Ean: "1234567890123", Sku: "SKU12345", BrutPrice: 1000, NetPrice: 850, AdditionalInfo: "Product A details", Quantity: 10},
	"PROD002": {Ean: "9876543210987", Sku: "SKU67890", BrutPrice: 2000, NetPrice: 1700, AdditionalInfo: "Product B details", Quantity: 5},
	"PROD003": {Ean: "1231231231231", Sku: "SKU54321", BrutPrice: 1500, NetPrice: 1300, AdditionalInfo: "Product C details", Quantity: 0}, // Out of stock
	"PROD004": {Ean: "3213213213213", Sku: "SKU09876", BrutPrice: 2500, NetPrice: 2200, AdditionalInfo: "Product D details", Quantity: 8},
	"PROD005": {Ean: "4564564564564", Sku: "SKU56789", BrutPrice: 3000, NetPrice: 2800, AdditionalInfo: "Product E details", Quantity: 2},
}

type Server struct {
	pb.UnimplementedProductsServiceServer
}

func (s *Server) CheckProductsAvailability(ctx context.Context, req *pb.ProductRequest) (*pb.ProductResponse, error) {
	fmt.Printf("Received request for productID: %s with quantity: %d\n", req.ProductId, req.Quantity)

	product, exists := fakeDB[req.ProductId]
	if !exists {
		return &pb.ProductResponse{
			Available:      false,
			AdditionalInfo: "Product not found",
		}, nil
	}

	if product.Quantity < req.Quantity {
		return &pb.ProductResponse{
			Available:      false,
			Ean:            product.Ean,
			Sku:            product.Sku,
			BrutPrice:      product.BrutPrice,
			NetPrice:       product.NetPrice,
			AdditionalInfo: "Insufficient stock",
		}, nil
	}

	return &pb.ProductResponse{
		Available:      true,
		Ean:            product.Ean,
		Sku:            product.Sku,
		BrutPrice:      product.BrutPrice,
		NetPrice:       product.NetPrice,
		AdditionalInfo: "Product is available",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterProductsServiceServer(grpcServer, &Server{})

	log.Println("gRPC server listening on port 50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
