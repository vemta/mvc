package usecase

import (
	"context"

	"github.com/vemta/common/entity"
	"github.com/vemta/mvc/internal/infra/repository"
	uow "github.com/vemta/mvc/pkg"
)

type FindCustomerOrdersUsecase struct {
	Uow uow.UowInterface
}

func NewFindCostomerOrdersUsecase(uow uow.UowInterface) *FindCustomerOrdersUsecase {
	return &FindCustomerOrdersUsecase{
		Uow: uow,
	}
}

func (u *FindCustomerOrdersUsecase) Execute(ctx context.Context, customer string) (*[]entity.Order, error) {
	return repository.GetCustomersRepository(ctx, u.Uow).FindCustomerOrders(ctx, customer)
}
