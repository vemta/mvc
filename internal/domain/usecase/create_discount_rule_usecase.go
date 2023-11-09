package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/vemta/common/entity"
	"github.com/vemta/mvc/internal/infra/repository"
	uow "github.com/vemta/mvc/pkg"
)

type CreateDiscountRuleUsecaseInput struct {
	ID                 string    `json:"ID"`
	Name               string    `json:"name"`
	DiscountRaw        float64   `json:"discount_raw"`
	DiscountPercentual float64   `json:"discount_percentual"`
	ApplyFirst         string    `json:"apply_first"` // raw | percentual
	AboveValue         float64   `json:"above_value"`
	BellowValue        float64   `json:"bellow_value"`
	ValidFrom          time.Time `json:"valid_from"`
	ValidUntil         time.Time `json:"valid_until"`
	AutoApply          bool      `json:"auto_apply"`
	Items              []string  `json:"item,omitempty"`
	Code               string    `json:"code,omitempty"`
	Type               string    `json:"type"`
}

type CreateDiscountRuleUsecase struct {
	Uow uow.UowInterface
}

func NewCreateDiscountRuleUsecase(uow uow.UowInterface) *CreateDiscountRuleUsecase {
	return &CreateDiscountRuleUsecase{
		Uow: uow,
	}
}

func (u *CreateDiscountRuleUsecase) Execute(ctx context.Context, rule CreateDiscountRuleUsecaseInput) error {

	discountRule := entity.DiscountRule{
		ID:                 rule.ID,
		Name:               rule.Name,
		DiscountRaw:        rule.DiscountRaw,
		DiscountPercentual: rule.DiscountPercentual,
		ApplyFirst:         rule.ApplyFirst,
		AboveValue:         rule.AboveValue,
		BellowValue:        rule.AboveValue,
		ValidFrom:          rule.ValidFrom,
		ValidUntil:         rule.ValidUntil,
		Code:               rule.Code,
		Type:               []rune(rule.Type)[0],
		AutoApply:          rule.AutoApply,
	}

	if []rune(rule.Type)[0] == 'O' {
		return repository.GetDiscountRulesRepository(ctx, u.Uow).CreateOrderDiscountRule(ctx, &entity.OrderDiscountRule{
			DiscountRule: &discountRule,
		})
	}

	if []rune(rule.Type)[0] == 'I' {
		return repository.GetDiscountRulesRepository(ctx, u.Uow).CreateItemDiscountRule(ctx, &entity.ItemDiscountRule{
			DiscountRule: &discountRule,
			Items:        rule.Items,
		})
	}

	return errors.New("discount code type not supported")
}
