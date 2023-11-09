package usecase

import (
	"context"
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/vemta/common/entity"
	"github.com/vemta/common/enum"
	"github.com/vemta/mvc/internal/infra/repository"
	uow "github.com/vemta/mvc/pkg"
)

type CreateOrderUsecaseInput struct {
	Customer      string               `json:"customer"`
	Items         *[]entity.OrderEntry `json:"items"`
	PaymentMethod int                  `json:"payment_method"`
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
		return errors.New("couldn't find user")
	}

	if len(*input.Items) <= 0 {
		return errors.New("cannot create empty order")
	}

	id := strings.Replace(uuid.New().String(), "-", "", -1)

	// TODO: Implement discount rules

	order := &entity.Order{
		ID:            id,
		Customer:      user,
		Items:         input.Items,
		PaymentMethod: *enum.GetPaymentMethod(input.PaymentMethod),
		Status:        enum.WaitingPaymentApproval,
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
