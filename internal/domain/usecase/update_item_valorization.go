package usecase

import (
	"context"
	"time"

	"github.com/vemta/common/entity"
	"github.com/vemta/mvc/internal/infra/repository"
	uow "github.com/vemta/mvc/pkg"
)

type ItemValorizationUpdateUsecaseInput struct {
	Item  string  `json:"item"`
	Cost  float64 `json:"cost"`
	Price float64 `json:"price"`
}

type ItemValuationUpdateUsecase struct {
	Uow uow.UowInterface
}

func NewItemValorizationUpdateUsecase(uow uow.UowInterface) *ItemValuationUpdateUsecase {
	return &ItemValuationUpdateUsecase{
		Uow: uow,
	}
}

func (u *ItemValuationUpdateUsecase) Execute(ctx context.Context, input ItemValorizationUpdateUsecaseInput) error {
	return u.Uow.Do(ctx, func(_ *uow.Uow) error {
		return repository.GetItemsRepository(ctx, u.Uow).UpdateItemValuation(ctx, input.Item, &entity.ItemValuation{
			LastCost:  input.Cost,
			LastPrice: input.Price,
			UpdatedAt: time.Now(),
		})
	})
}
