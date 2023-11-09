package usecase

import (
	"context"
	"time"

	"github.com/vemta/mvc/internal/infra/repository"
	uow "github.com/vemta/mvc/pkg"
)

type FindActiveDiscountRulesUsecase struct {
	Uow uow.UowInterface
}

func NewFindActiveDiscountRulesUsecase(uow uow.UowInterface) *FindCustomerUsecase {
	return &FindCustomerUsecase{
		Uow: uow,
	}
}

func (u *FindActiveDiscountRulesUsecase) Execute(ctx context.Context, customer string) (*[]any, error) {
	return repository.GetDiscountRulesRepository(ctx, u.Uow).FindActiveDiscountRules(ctx, time.Now())
}
