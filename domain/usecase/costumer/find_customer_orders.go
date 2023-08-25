package usecase

import (
	"context"
	"errors"

	uow "github.com/vemta/mvc/pkg"
)

type FindCustomerOrdersUsecase struct {
	Uow uow.UowInterface
}

func NewFindCostomerOrdersUsecase(uow uow.UowInterface) *FindCustomerOrdersUsecase {
	return &FindCustomerOrdersUsecase{
		Uow: uow,
	}
}

func (u *FindCustomerOrdersUsecase) Execute(ctx context.Context, customer string) (float64, error) {
	return 0, errors.New("not implemented yet")
}
