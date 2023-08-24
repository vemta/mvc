package usecase

import (
	"context"

	"github.com/vemta/mvc/domain/repository"
	uow "github.com/vemta/mvc/pkg"
)

type FindCustomerOrdersUsecaseInput struct {
	Customer string `json:"costumer"`
}

type FindCustomerOrdersUsecase struct {
	Uow uow.UowInterface
}

func NewFindCostomerOrdersUsecase(uow uow.UowInterface) *FindCustomerOrdersUsecase {
	return &FindCustomerOrdersUsecase{
		Uow: uow,
	}
}

func (u *FindCustomerOrdersUsecase) Execute(ctx context.Context, input FindCustomerOrdersUsecaseInput) (float64, error) {
	item, err := u.getItemRepository(ctx).FindItem(ctx, input.ID)
	if err != nil {
		return 0, err
	}

	valuation := item.Valuation
	return (valuation.LastPrice - valuation.DiscountRaw) * (1 - valuation.DiscountPercentual), nil
}

func (u *FindCustomerOrdersUsecase) getItemRepository(ctx context.Context) repository.ItemRepositoryInterface {
	itemRepository, err := u.Uow.GetRepository(ctx, "ItemRepository")
	if err != nil {
		panic(err)
	}
	return itemRepository.(repository.ItemRepositoryInterface)
}
