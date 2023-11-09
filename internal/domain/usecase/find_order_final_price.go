package usecase

import (
	"context"
	"math"
	"time"

	"github.com/vemta/mvc/internal/infra/repository"
	uow "github.com/vemta/mvc/pkg"
)

type FindOrderFinalPriceUsecaseInput struct {
	DiscountCodes []string `json:"discount_codes"`
	Items         []string `json:"items"`
}

type FindOrderFinalPriceUsecase struct {
	Uow uow.UowInterface
}

func NewFindOrderFinalPriceUsecase(uow uow.UowInterface) *FindOrderFinalPriceUsecase {
	return &FindOrderFinalPriceUsecase{
		Uow: uow,
	}
}

func (u *FindOrderFinalPriceUsecase) Execute(ctx context.Context, input FindOrderFinalPriceUsecaseInput) (float64, error) {

	current := 0.0
	for _, detail := range input.Items {
		uc := NewFindItemFinalPriceUsecase(u.Uow)
		price, err := uc.Execute(ctx, detail)
		if err != nil {
			return 0, err
		}
		current += price
	}

	autoApply, err := repository.GetDiscountRulesRepository(ctx, u.Uow).FindAutoApplyDiscountRulesForOrder(ctx, current, time.Now())
	if err != nil {
		return current, nil
	}

	for _, discount := range *autoApply {
		if discount.ApplyFirst == "raw" {
			current -= discount.DiscountRaw
		} else {
			current *= (1 - discount.DiscountPercentual)
		}
	}

	return math.Max(0, current), nil
}
