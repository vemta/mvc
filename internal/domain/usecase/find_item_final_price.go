package usecase

import (
	"context"
	"math"

	"github.com/vemta/mvc/internal/infra/repository"
	uow "github.com/vemta/mvc/pkg"
)

type FindItemFinalPriceUsecase struct {
	Uow uow.UowInterface
}

func NewFindItemFinalPriceUsecase(uow uow.UowInterface) *FindItemFinalPriceUsecase {
	return &FindItemFinalPriceUsecase{
		Uow: uow,
	}
}

func (u *FindItemFinalPriceUsecase) Execute(ctx context.Context, id string) (float64, error) {
	item, err := repository.GetItemsRepository(ctx, u.Uow).FindItem(ctx, id)
	if err != nil {
		return 0, err
	}

	availableDiscounts, err := repository.GetDiscountRulesRepository(ctx, u.Uow).FindValidDiscountRulesForItem(ctx, id)
	if err != nil {
		return item.Valuation.LastPrice, nil
	}

	current := item.Valuation.LastPrice
	for _, discount := range *availableDiscounts {
		if discount.ApplyFirst == "RAW" {
			current = (current - discount.DiscountRaw) * (1 - discount.DiscountPercentual)
		} else {
			current = (current * (1 - discount.DiscountPercentual)) - discount.DiscountRaw
		}
	}

	return math.Max(0, current), nil
}
