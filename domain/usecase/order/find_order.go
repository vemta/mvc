package usecase

import (
	"context"

	"github.com/vemta/common/entity"
	"github.com/vemta/mvc/internal/infra/repository"
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
	return repository.GetOrdersRepository(ctx, u.Uow).FindOrder(ctx, input.ID)
}
