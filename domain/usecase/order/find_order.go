package usecase

import (
	"context"

	"github.com/vemta/common/entity"
	"github.com/vemta/mvc/domain/repository"
	uow "github.com/vemta/mvc/pkg"
)

type FindOrderUsecaseInput struct {
	ID string `json:"id"`
}

type FindOrderUsecase struct {
	Uow uow.UowInterface
}

func NewFindOrderUsecase(uow uow.UowInterface) *FindOrderUsecase {
	return &FindOrderUsecase{
		Uow: uow,
	}
}

func (u *FindOrderUsecase) Execute(ctx context.Context, input FindOrderUsecaseInput) (*entity.Order, error) {
	return u.getOrderRepository(ctx).FindOrder(ctx, input.ID)
}

func (u *FindOrderUsecase) getOrderRepository(ctx context.Context) repository.OrderRepositoryInterface {
	repo, err := u.Uow.GetRepository(ctx, "OrderRepository")
	if err != nil {
		panic(err)
	}
	return repo.(repository.OrderRepositoryInterface)
}
