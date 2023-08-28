package usecase

import (
	"context"

	"github.com/vemta/common/entity"
	"github.com/vemta/mvc/internal/infra/repository"
	uow "github.com/vemta/mvc/pkg"
)

type CreateOrderUsecaseInput struct {
	ID                 string               `json:"id"`
	Customer           string               `json:"customer"`
	Items              *[]entity.OrderEntry `json:"items"`
	PaymentMethod      int                  `json:"payment_method"`
	DiscountRaw        float64              `json:"discount_raw"`
	DiscountPercentual float64              `json:"discount_percentual"`
	Status             uint8                `json:"status"`
}

type CreateOrderUsecase struct {
	Uow uow.UowInterface
}

func NewCreateOrderUsecase(uow uow.UowInterface) *CreateOrderUsecase {
	return &CreateOrderUsecase{
		Uow: uow,
	}
}

func (u *CreateOrderUsecase) Execute(ctx context.Context, input CreateOrderUsecaseInput) error {

	user, err := repository.GetCustomersRepository(ctx, uow.GetCurrent()).FindCustomer(ctx, input.Customer)
	if err != nil {
		return err
	}

	order := &entity.Order{
		ID:                 input.ID,
		Customer:           user,
		Items:              input.Items,
		PaymentMethod:      input.PaymentMethod,
		DiscountRaw:        input.DiscountRaw,
		DiscountPercentual: input.DiscountPercentual,
		Status:             input.Status,
	}

	current := 0.0
	for _, detail := range *order.Items {
		uc := NewFindItemFinalPriceUsecase(u.Uow)
		price, err := uc.Execute(ctx, detail.Item.ID)
		if err != nil {
			return err
		}
		current += price * float64(detail.Quantity)
	}
	current -= order.DiscountRaw
	current *= (1 - order.DiscountPercentual)
	order.Price = current

	return repository.GetOrdersRepository(ctx, u.Uow).Create(ctx, order)
}
