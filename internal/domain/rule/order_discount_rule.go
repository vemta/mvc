package rule

import (
	"time"

	"github.com/vemta/common/entity"
)

type OrderDiscountRule struct {
	ID                 string
	Name               string
	DiscountRaw        float64
	DiscountPercentual float64
	ApplyFirst         string // raw | percentual
	AboveValue         float64
	BellowValue        float64
	ValidFrom          time.Time
	ValidUntil         time.Time
}

func (d *OrderDiscountRule) TryApply(order *entity.Order) (bool, float64) {

	if order.Price < d.AboveValue && d.AboveValue != -1 {
		return false, order.Price
	}

	if order.Price > d.BellowValue && d.BellowValue != -1 {
		return false, order.Price
	}

	if !d.ValidFrom.IsZero() && !time.Now().After(d.ValidFrom) {
		return false, order.Price
	}

	if !d.ValidUntil.IsZero() && !time.Now().Before(d.ValidUntil) {
		return false, order.Price
	}

	if d.ApplyFirst == "raw" {
		return true, (order.Price - d.DiscountRaw) * (1 - d.DiscountPercentual)
	}

	return true, (order.Price * (1 - d.DiscountPercentual)) - d.DiscountRaw
}
