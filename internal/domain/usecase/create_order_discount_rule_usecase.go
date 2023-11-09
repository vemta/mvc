package usecase

import (
	"context"
	"time"

	"github.com/vemta/common/entity"
	"github.com/vemta/mvc/internal/infra/repository"
	uow "github.com/vemta/mvc/pkg"
)

type CreateOrderDiscountRuleUsecaseInput struct {
	ID                 string    `json:"ID"`
	Name               string    `json:"name"`
	DiscountRaw        float64   `json:"discount_raw"`
	DiscountPercentual float64   `json:"discount_percentual"`
	ApplyFirst         string    `json:"apply_first"` // raw | percentual
	AboveValue         float64   `json:"above_value"`
	BellowValue        float64   `json:"bellow_value"`
	ValidFrom          time.Time `json:"valid_from"`
	ValidUntil         time.Time `json:"valid_until"`
	Items              []string  `json:"item"`
}

type CreateOrderDiscountRuleUsecase struct {
	Uow uow.UowInterface
}

func NewCreateOrderDiscountRuleUsecase(uow uow.UowInterface) *CreateOrderDiscountRuleUsecase {
	return &CreateOrderDiscountRuleUsecase{
		Uow: uow,
	}
}

func (u *CreateOrderDiscountRuleUsecase) Execute(ctx context.Context, rule CreateOrderDiscountRuleUsecaseInput) error {

	discountRule := entity.OrderDiscountRule{
		ID:                 rule.ID,
		Name:               rule.Name,
		DiscountRaw:        rule.DiscountRaw,
		DiscountPercentual: rule.DiscountPercentual,
		ApplyFirst:         rule.ApplyFirst,
		AboveValue:         rule.AboveValue,
		BellowValue:        rule.AboveValue,
		ValidFrom:          rule.ValidFrom,
		ValidUntil:         rule.ValidUntil,
	}
	return repository.GetDiscountRulesRepository(ctx, uow.GetCurrent()).CreateOrderDiscountRule(ctx, &discountRule)

}
