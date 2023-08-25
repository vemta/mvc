package grpc

import (
	"context"

	usecase "github.com/vemta/mvc/domain/usecase/order"
	"github.com/vemta/mvc/internal/infra/db"
	"github.com/vemta/mvc/internal/infra/grpc/pb"
	uow "github.com/vemta/mvc/pkg"
)

type orderServiceServer struct {
	pb.UnimplementedOrderServiceServer
	uow     *uow.UowInterface
	queries *db.Queries
}

func NewOrderService(uow *uow.UowInterface, queries *db.Queries) *orderServiceServer {
	return &orderServiceServer{
		uow:     uow,
		queries: queries,
	}
}

func (s *orderServiceServer) GetFinalPrice(ctx context.Context, input *pb.GetFinalPriceRequest) (*pb.GetFinalPriceResponse, error) {
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
