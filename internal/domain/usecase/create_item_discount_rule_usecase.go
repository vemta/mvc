package usecase

import (
	"context"
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/vemta/common/entity"
	"github.com/vemta/mvc/internal/infra/repository"
	uow "github.com/vemta/mvc/pkg"
)

type CreateItemDiscountRuleUsecaseInput struct {
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

type CreateItemDiscountRuleUsecase struct {
	Uow uow.UowInterface
}

func NewCreateItemDiscountRuleUsecase(uow uow.UowInterface) *CreateItemDiscountRuleUsecase {
	return &CreateItemDiscountRuleUsecase{
		Uow: uow,
	}
}

func (u *CreateItemDiscountRuleUsecase) Execute(ctx context.Context, rule CreateItemDiscountRuleUsecaseInput) error {

	id, err := gonanoid.Generate("abcdefghijklmnopqrstuvwxyz0123456789", 12)

	if err != nil {
		return err
	}

	discountRule := entity.ItemDiscountRule{
		ID:                 id,
		Name:               rule.Name,
		DiscountRaw:        rule.DiscountRaw,
		DiscountPercentual: rule.DiscountPercentual,
		ApplyFirst:         rule.ApplyFirst,
		AboveValue:         rule.AboveValue,
		BellowValue:        rule.AboveValue,
		ValidFrom:          rule.ValidFrom,
		ValidUntil:         rule.ValidUntil,
		Items:              rule.Items,
	}
	return repository.GetDiscountRulesRepository(ctx, uow.GetCurrent()).CreateItemDiscountRule(ctx, &discountRule)

}
