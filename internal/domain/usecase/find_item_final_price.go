package usecase

import (
	"context"
	"time"

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

	currentPrice := item.Valuation.LastPrice
	rules, e := repository.GetDiscountRulesRepository(ctx, u.Uow).FindAutoApplyDiscountRulesForItem(ctx, id, time.Now())
	if e != nil {
		return currentPrice, nil
	}

	for _, rule := range *rules {
		if rule.ApplyFirst == "raw" {
			currentPrice = (currentPrice - rule.DiscountRaw) * (1 - rule.DiscountPercentual)
		} else {
			currentPrice = (currentPrice * (1 - rule.DiscountPercentual)) - rule.DiscountRaw
		}
	}

	return currentPrice, nil
}
