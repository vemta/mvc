package usecase

import (
	"context"

	"github.com/vemta/common/entity"
	"github.com/vemta/mvc/internal/infra/repository"
	uow "github.com/vemta/mvc/pkg"
)

type FindCustomerUsecase struct {
	Uow uow.UowInterface
}

func NewFindCustomerUsecase(uow uow.UowInterface) *FindCustomerUsecase {
	return &FindCustomerUsecase{
		Uow: uow,
	}
}

func (u *FindCustomerUsecase) Execute(ctx context.Context, customer string) (*entity.Customer, error) {
	return repository.GetCustomersRepository(ctx, u.Uow).FindCustomer(ctx, customer)
}
