package usecase

import (
	"context"
	"errors"

	uow "github.com/vemta/mvc/pkg"
)

type FindCustomerUsecaseInput struct {
	Customer string `json:"costumer"`
}

type FindCustomerUsecase struct {
	Uow uow.UowInterface
}

func NewFindCustomerUsecase(uow uow.UowInterface) *FindCustomerUsecase {
	return &FindCustomerUsecase{
		Uow: uow,
	}
}

func (u *FindCustomerUsecase) Execute(ctx context.Context, input FindCustomerUsecaseInput) (float64, error) {
	return 0, errors.New("not implemented yet")
}
