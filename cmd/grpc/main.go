package main

import (
	"log"
	"net"

	"github.com/diegobcaetano/product-service/internal/api/config"
	"github.com/diegobcaetano/product-service/internal/logging"
	pb "github.com/diegobcaetano/product-service/pkg/domain/grpc"
	"github.com/diegobcaetano/product-service/pkg/handler"
	repository "github.com/diegobcaetano/product-service/pkg/infrastructure/db"
	"github.com/diegobcaetano/product-service/pkg/usecase"
	"google.golang.org/grpc"
)


func main() {

	logger := logging.NewSlogAdapter()
	env := config.NewEnvLoad(logger).Load("../../")
	dbSession, err := repository.NewCassandraSession(env.GetConfig())
	defer dbSession.Close()
	if err != nil {
		panic(err)
	}
	productRepository := repository.NewCassandraProductRepository(dbSession, logger)
	var useCase usecase.ProductUseCase = usecase.NewProductService(productRepository)

	// Inicialize o servidor gRPC
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterProductServiceServer(s, &handler.ProductGrpcServer{
		UseCase: useCase,
		Logger: logger,
	})
	log.Println("Server is running on port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
