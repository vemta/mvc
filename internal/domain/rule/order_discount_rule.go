package rule

import (
	"time"

	"github.com/vemta/common/entity"
)

type OrderDiscountRuleParams struct {
	AboveValue  float64
	BellowValue float64
	AfterDate   time.Time
	BeforeDate  time.Time
}

type OrderDiscountRule struct {
	ID                 string
	Name               string
	DiscountRaw        float64
	DiscountPercentual float64
	ApplyFirst         string // raw | percentual
	Params             OrderDiscountRuleParams
}

func (d *OrderDiscountRule) TryApply(order *entity.Order) (bool, float64) {

	if order.Price < d.Params.AboveValue && d.Params.AboveValue != -1 {
		return false, order.Price
	}

	if order.Price > d.Params.BellowValue && d.Params.BellowValue != -1 {
		return false, order.Price
	}

	if !d.Params.AfterDate.IsZero() && !time.Now().After(d.Params.AfterDate) {
		return false, order.Price
	}

	if !d.Params.BeforeDate.IsZero() && !time.Now().Before(d.Params.BeforeDate) {
		return false, order.Price
	}

	if d.ApplyFirst == "raw" {
		return true, (order.Price - d.DiscountRaw) * (1 - d.DiscountPercentual)
	}

	return true, (order.Price * (1 - d.DiscountPercentual)) - d.DiscountRaw
}
