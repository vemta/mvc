package usecase

import (
	"context"
	"time"

	"github.com/vemta/common/entity"
	"github.com/vemta/mvc/internal/infra/repository"
	uow "github.com/vemta/mvc/pkg"
)

type CreateOrderDiscountRuleUsecaseInput struct {
	ID                 string
	Name               string
	Items              []string
	DiscountRaw        float64
	DiscountPercentual float64
	ApplyFirst         string // raw | percentual
	AboveValue         float64
	BellowValue        float64
	ValidFrom          time.Time
	ValidUntil         time.Time
}

type CreateOrderDiscountRuleUsecase struct {
	Uow uow.UowInterface
}

func NewCreateOrderDiscountRuleUsecase(uow uow.UowInterface) *CreateOrderDiscountRuleUsecase {
	return &CreateOrderDiscountRuleUsecase{
		Uow: uow,
	}
}

func (u *CreateOrderDiscountRuleUsecase) Execute(ctx context.Context, input CreateOrderDiscountRuleUsecaseInput) error {

	return repository.GetDiscountRulesRepository(ctx, uow.GetCurrent()).CreateOrderDiscountRule(ctx, &entity.OrderDiscountRule{
		ID:                 input.ID,
		Name:               input.Name,
		DiscountRaw:        input.DiscountRaw,
		DiscountPercentual: input.DiscountPercentual,
		ApplyFirst:         input.ApplyFirst,
		AboveValue:         input.AboveValue,
		BellowValue:        input.BellowValue,
		ValidFrom:          input.ValidFrom,
		ValidUntil:         input.ValidUntil,
	})

}
