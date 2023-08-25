package usecase

import (
	"context"
	"math"

	usecase "github.com/vemta/mvc/domain/usecase/item"
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

func (u *FindOrderFinalPriceUsecase) Execute(ctx context.Context, id string) (float64, error) {
	order, err := repository.GetOrdersRepository(ctx, u.Uow).FindOrder(ctx, id)
	if err != nil {
		return 0, err
	}

	current := 0.0
	for _, detail := range *order.Items {
		uc := usecase.NewFindItemFinalPriceUsecase(u.Uow)
		price, err := uc.Execute(ctx, detail.Item.ID)
		if err != nil {
			return 0, err
		}
		current += price
	}

	current -= current - order.DiscountRaw
	current *= order.DiscountPercentual

	return math.Max(0, current), nil
}
