package usecase

import (
	"context"

	"github.com/vemta/common/entity"
	"github.com/vemta/mvc/internal/infra/repository"
	uow "github.com/vemta/mvc/pkg"
)

type FindOrderUsecase struct {
	Uow uow.UowInterface
}

func NewFindOrderUsecase(uow uow.UowInterface) *FindOrderUsecase {
	return &FindOrderUsecase{
		Uow: uow,
	}
}

func (u *FindOrderUsecase) Execute(ctx context.Context, id string) (*entity.Order, error) {
	return repository.GetOrdersRepository(ctx, u.Uow).FindOrder(ctx, id)
}
