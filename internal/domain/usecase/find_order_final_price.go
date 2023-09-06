package usecase

import (
	"context"
	"math"

	"github.com/vemta/mvc/internal/infra/repository"
	uow "github.com/vemta/mvc/pkg"
)

type FindOrderFinalPriceUsecase struct {
	Uow uow.UowInterface
}

func NewFindOrderFinalPriceUsecase(uow uow.UowInterface) *FindOrderFinalPriceUsecase {
	return &FindOrderFinalPriceUsecase{
		Uow: uow,
	}
}

func (u *FindOrderFinalPriceUsecase) Execute(ctx context.Context, items []string) (float64, error) {
	current := 0.0
	for _, item := range items {
		uc := NewFindItemFinalPriceUsecase(u.Uow)
		price, err := uc.Execute(ctx, item)
		if err != nil {
			return 0, err
		}
		current += price
	}

	availableDiscounts, err := repository.GetDiscountRulesRepository(ctx, u.Uow).FindValidOrderDiscountRules(ctx)
	if err != nil {
		return 0, err
	}
	for _, discount := range *availableDiscounts {
		if discount.ApplyFirst == "RAW" {
			current = (current - discount.DiscountRaw) * (1 - discount.DiscountPercentual)
		} else {
			current = (current * (1 - discount.DiscountPercentual)) - discount.DiscountRaw
		}
	}

	return math.Max(0, current), nil
}
