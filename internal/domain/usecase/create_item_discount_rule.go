package usecase

import (
	"context"
	"time"

	"github.com/vemta/common/entity"
	"github.com/vemta/mvc/internal/infra/repository"
	uow "github.com/vemta/mvc/pkg"
)

type CreateItemDiscountRuleUsecaseInput struct {
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

type CreateItemDiscountRuleUsecase struct {
	Uow uow.UowInterface
}

func NewCreateItemDiscountRuleUsecase(uow uow.UowInterface) *CreateItemDiscountRuleUsecase {
	return &CreateItemDiscountRuleUsecase{
		Uow: uow,
	}
}

func (u *CreateItemDiscountRuleUsecase) Execute(ctx context.Context, input CreateItemDiscountRuleUsecaseInput) error {

	return repository.GetDiscountRulesRepository(ctx, uow.GetCurrent()).CreateItemDiscountRule(ctx, &entity.ItemDiscountRule{
		ID:                 input.ID,
		Name:               input.Name,
		Items:              input.Items,
		DiscountRaw:        input.DiscountRaw,
		DiscountPercentual: input.DiscountPercentual,
		ApplyFirst:         input.ApplyFirst,
		AboveValue:         input.AboveValue,
		BellowValue:        input.BellowValue,
		ValidFrom:          input.ValidFrom,
		ValidUntil:         input.ValidUntil,
	})

}
