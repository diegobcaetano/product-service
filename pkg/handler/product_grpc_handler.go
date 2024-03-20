package handler

import (
	"context"

	"github.com/diegobcaetano/product-service/internal/logging"
	pb "github.com/diegobcaetano/product-service/pkg/domain/grpc"
	model "github.com/diegobcaetano/product-service/pkg/domain/model/product"
	"github.com/diegobcaetano/product-service/pkg/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductGrpcServer struct{
	UseCase usecase.ProductUseCase
	Logger  logging.Logger
	pb.UnimplementedProductServiceServer
}

func (s *ProductGrpcServer) GetProductByID(ctx context.Context, req *pb.ProductByIDRequest) (*pb.ProductResponse, error) {

	p, err := s.UseCase.GetProductByID(req.GetProductId())
	if err != nil {
		s.Logger.Error("Something went wrong while fetching the product", "error", err.Error())
		return &pb.ProductResponse{
			Product: &pb.Product{},
		}, nil
	}
	
	protoProduct, err := model.ProductToProto(p)

	if err != nil {
		s.Logger.Error("Something went wrong while converting the product to Protobuf", "error", err.Error())
	}

	return &pb.ProductResponse{
		Product: protoProduct,
	}, nil
}

func (s *ProductGrpcServer) ProductHasStock(ctx context.Context, req *pb.ProductHasStockRequest) (*pb.ProductHasStockResponse, error) {

	result, err := s.UseCase.ProductHasStock(req.GetProductId(),req.GetSellerId())
	if err != nil {
		s.Logger.Error("Something went wrong while fetching the product", "error", err.Error())
		return nil, status.Errorf(codes.Internal, "Something went wrong while fetching the product stock: %v", err)
	}

	return &pb.ProductHasStockResponse{
		Result: result,
	}, nil
}
