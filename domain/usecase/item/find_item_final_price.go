package usecase

import (
	"context"

	"github.com/vemta/mvc/domain/repository"
	uow "github.com/vemta/mvc/pkg"
)

type FindItemFinalPriceUsecaseInput struct {
	ID string `json:"id"`
}

type FindItemFinalPriceUsecase struct {
	Uow uow.UowInterface
}

func NewFindItemFinalPriceUsecase(uow uow.UowInterface) *FindItemFinalPriceUsecase {
	return &FindItemFinalPriceUsecase{
		Uow: uow,
	}
}

func (u *FindItemFinalPriceUsecase) Execute(ctx context.Context, input FindItemFinalPriceUsecaseInput) (float64, error) {
	item, err := u.getItemRepository(ctx).FindItem(ctx, input.ID)
	if err != nil {
		return 0, err
	}

	valuation := item.Valuation
	return (valuation.LastPrice - valuation.DiscountRaw) * (1 - valuation.DiscountPercentual), nil
}

func (u *FindItemFinalPriceUsecase) getItemRepository(ctx context.Context) repository.ItemRepositoryInterface {
	itemRepository, err := u.Uow.GetRepository(ctx, "ItemRepository")
	if err != nil {
		panic(err)
	}
	return itemRepository.(repository.ItemRepositoryInterface)
}
