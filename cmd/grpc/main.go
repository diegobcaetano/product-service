package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/diegobcaetano/product-service/pkg/domain/grpc"
	"google.golang.org/grpc"
)

// server implementa grpc.ProductServiceServer
type server struct{
	pb.UnimplementedProductServiceServer
}

func (s *server) GetProductByID(ctx context.Context, req *pb.ProductByIDRequest) (*pb.ProductResponse, error) {
	// Lógica para obter o produto com o ID fornecido
	// Aqui você pode chamar sua lógica de negócios ou acessar o banco de dados, etc.

	fmt.Printf("Server info %s", req.ProductId)
	// Construa e retorne a resposta
	return &pb.ProductResponse{
		Product: &pb.Product{
			Id: req.ProductId,
		},
	}, nil
}

func main() {
	// Inicialize o servidor gRPC
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProductServiceServer(s, &server{})
	log.Println("Server is running on port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
