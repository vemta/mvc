package grpc

import (
	"context"
	"github.com/vemta/mvc/internal/domain/usecase"

	"github.com/vemta/mvc/internal/infra/db"
	"github.com/vemta/mvc/internal/infra/grpc/pb"
	uow "github.com/vemta/mvc/pkg"
)

type OrderServiceServer struct {
	pb.UnimplementedOrderServiceServer
	uow     *uow.UowInterface
	queries *db.Queries
}

func NewOrderService(uow *uow.UowInterface, queries *db.Queries) *OrderServiceServer {
	return &OrderServiceServer{
		uow:     uow,
		queries: queries,
	}
}

func (s *OrderServiceServer) GetFinalPrice(ctx context.Context, input *pb.GetFinalPriceRequest) (*pb.GetFinalPriceResponse, error) {
	uc := usecase.NewFindOrderFinalPriceUsecase(*s.uow)
	finalPrice, err := uc.Execute(ctx, input.Order)
	if err != nil {
		return nil, err
	}
	return &pb.GetFinalPriceResponse{
		Success: true,
		Price:   finalPrice,
	}, nil
}
