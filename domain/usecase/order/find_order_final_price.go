package usecase

import (
	"context"
	"math"

	"github.com/vemta/mvc/domain/repository"
	usecase "github.com/vemta/mvc/domain/usecase/item"
	uow "github.com/vemta/mvc/pkg"
)

type FindOrderFinalPriceUsecaseInput struct {
	ID string `json:"id"`
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
	order, err := u.getOrderRepository(ctx).FindOrder(ctx, input.ID)
	if err != nil {
		return 0, err
	}

	current := 0.0
	for _, detail := range *order.Items {
		uc := usecase.NewFindItemFinalPriceUsecase(u.Uow)
		price, err := uc.Execute(ctx, usecase.FindItemFinalPriceUsecaseInput{
			ID: detail.Item.ID,
		})
		if err != nil {
			return 0, err
		}
		current += price
	}

	current -= current - order.DiscountRaw
	current *= order.DiscountPercentual

	return math.Max(0, current), nil

}

func (u *FindOrderFinalPriceUsecase) getOrderRepository(ctx context.Context) repository.OrderRepositoryInterface {
	repo, err := u.Uow.GetRepository(ctx, "OrderRepository")
	if err != nil {
		panic(err)
	}
	return repo.(repository.OrderRepositoryInterface)
}
