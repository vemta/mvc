package usecase

import (
	"context"

	"github.com/vemta/common/enum/orderstatus"
	"github.com/vemta/mvc/internal/infra/repository"
	uow "github.com/vemta/mvc/pkg"
)

type UpdateOrderStatusUsecaseInput struct {
	Order  string                  `json:"order"`
	Status orderstatus.OrderStatus `json:"status"`
}

type UpdateOrderStatusUsecase struct {
	Uow uow.UowInterface
}

func NewUpdateOrderStatusUsecase(uow uow.UowInterface) *UpdateOrderStatusUsecase {
	return &UpdateOrderStatusUsecase{
		Uow: uow,
	}
}

func (u *UpdateOrderStatusUsecase) Execute(ctx context.Context, input UpdateOrderStatusUsecaseInput) error {

	return repository.GetOrdersRepository(ctx, uow.GetCurrent()).UpdateOrderStatus(ctx, input.Order, int(input.Status))

}
