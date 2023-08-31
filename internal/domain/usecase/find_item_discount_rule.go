package usecase

import (
	"context"

	"github.com/vemta/mvc/internal/infra/repository"
	uow "github.com/vemta/mvc/pkg"
)

type FindItemDiscountRuleUsecase struct {
	Uow uow.UowInterface
}

func NewFindItemDiscountRuleUsecase(uow uow.UowInterface) *FindItemDiscountRuleUsecase {
	return &FindItemDiscountRuleUsecase{
		Uow: uow,
	}
}

func (u *FindItemDiscountRuleUsecase) Execute(ctx context.Context, id string) (float64, error) {
	item, err := repository.GetItemsRepository(ctx, u.Uow).FindItem(ctx, id)
	if err != nil {
		return 0, err
	}

	valuation := item.Valuation
	return (valuation.LastPrice - valuation.DiscountRaw) * (1 - valuation.DiscountPercentual), nil
}
