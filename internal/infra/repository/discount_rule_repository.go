package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/vemta/common/entity"
	"github.com/vemta/mvc/internal/infra/db"
)

type DiscountRuleRepository struct {
	dbConn *sql.DB
	*db.Queries
	Repository
}

func NewDiscountRuleRepository(dbConn *sql.DB) *DiscountRuleRepository {
	return &DiscountRuleRepository{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

func (r *DiscountRuleRepository) FindValidDiscountRulesForItem(ctx context.Context, item string, price float64, time time.Time) (*[]entity.ItemDiscountRule, error) {
	dr, err := r.Queries.FindValidDiscountRulesForItem(ctx, db.FindValidDiscountRulesForItemParams{
		Item:        item,
		Validfrom:   time,
		Validuntil:  time,
		Abovevalue:  price,
		Bellowvalue: price,
	})

	if err != nil {
		return nil, err
	}

	discountRules := make([]entity.ItemDiscountRule, len(dr))

	for _, rule := range dr {
		discountRules = append(discountRules, entity.ItemDiscountRule{
			ID:                 rule.ID,
			Name:               rule.Name,
			Items:              []string{rule.Item},
			DiscountRaw:        rule.Discountraw,
			DiscountPercentual: rule.Discountpercentual,
			ApplyFirst:         rule.Applyfirst,
			AboveValue:         rule.Abovevalue,
			BellowValue:        rule.Abovevalue,
			ValidFrom:          rule.Validfrom,
			ValidUntil:         rule.Validuntil,
		})
	}

	return &discountRules, nil
}

func (r *DiscountRuleRepository) FindValidDiscountRulesForOrder(ctx context.Context, price float64, time time.Time) (*[]entity.OrderDiscountRule, error) {
	dr, err := r.Queries.FindValidDiscountRulesForOrder(ctx, db.FindValidDiscountRulesForOrderParams{
		Validfrom:   time,
		Validuntil:  time,
		Abovevalue:  price,
		Bellowvalue: price,
	})

	if err != nil {
		return nil, err
	}

	discountRules := make([]entity.OrderDiscountRule, len(dr))

	for _, rule := range dr {
		discountRules = append(discountRules, entity.OrderDiscountRule{
			ID:                 rule.ID,
			Name:               rule.Name,
			DiscountRaw:        rule.Discountraw,
			DiscountPercentual: rule.Discountpercentual,
			ApplyFirst:         rule.Applyfirst,
			AboveValue:         rule.Abovevalue,
			BellowValue:        rule.Abovevalue,
			ValidFrom:          rule.Validfrom,
			ValidUntil:         rule.Validuntil,
		})
	}

	return &discountRules, nil
}

func (r *DiscountRuleRepository) CreateItemDiscountRule(ctx context.Context, rule *entity.ItemDiscountRule) error {
	if err := r.Queries.CreateDiscountRule(ctx, db.CreateDiscountRuleParams{
		ID:                 rule.ID,
		Name:               rule.Name,
		Discountraw:        rule.DiscountRaw,
		Discountpercentual: rule.DiscountPercentual,
		Applyfirst:         rule.ApplyFirst,
		Abovevalue:         rule.AboveValue,
		Bellowvalue:        rule.BellowValue,
		Validfrom:          rule.ValidFrom,
		Validuntil:         rule.ValidUntil,
		Type:               "I",
	}); err != nil {
		return err
	}
	for _, item := range rule.Items {
		if err := r.Queries.CreateItemForDiscountRule(ctx, db.CreateItemForDiscountRuleParams{
			Discountrule: rule.ID,
			Item:         item,
		}); err != nil {
			return err
		}
	}
	return nil
}

func (r *DiscountRuleRepository) CreateOrderDiscountRule(ctx context.Context, rule *entity.OrderDiscountRule) error {
	if err := r.Queries.CreateDiscountRule(ctx, db.CreateDiscountRuleParams{
		ID:                 rule.ID,
		Name:               rule.Name,
		Discountraw:        rule.DiscountRaw,
		Discountpercentual: rule.DiscountPercentual,
		Applyfirst:         rule.ApplyFirst,
		Abovevalue:         rule.AboveValue,
		Bellowvalue:        rule.BellowValue,
		Validfrom:          rule.ValidFrom,
		Validuntil:         rule.ValidUntil,
		Type:               "O",
	}); err != nil {
		return err
	}
	return nil
}

func (r *DiscountRuleRepository) FindActiveDiscountRules(ctx context.Context, time time.Time) (*[]entity.DiscountRule, error) {

	rules, err := r.Queries.FindActiveDiscountRules(ctx, db.FindActiveDiscountRulesParams{
		Validfrom:  time,
		Validuntil: time,
	})

	if err != nil {
		return nil, err
	}

	discountRules := make([]entity.DiscountRule, 0)

	for _, rule := range rules {
		if rule.Type == "I" {
			items, err := r.FindItemsForDiscountRule(ctx, rule.ID)
			if err != nil {
				return nil, err
			}
			discountRules = append(discountRules, entity.ItemDiscountRule{
				ID:                 rule.ID,
				Name:               rule.Name,
				DiscountRaw:        rule.Discountraw,
				DiscountPercentual: rule.Discountpercentual,
				ApplyFirst:         rule.Applyfirst,
				AboveValue:         rule.Abovevalue,
				BellowValue:        rule.Bellowvalue,
				ValidFrom:          rule.Validfrom,
				ValidUntil:         rule.Validuntil,
				Items:              items,
			})
		} else {
			discountRules = append(discountRules, entity.OrderDiscountRule{
				ID:                 rule.ID,
				Name:               rule.Name,
				DiscountRaw:        rule.Discountraw,
				DiscountPercentual: rule.Discountpercentual,
				ApplyFirst:         rule.Applyfirst,
				AboveValue:         rule.Abovevalue,
				BellowValue:        rule.Bellowvalue,
				ValidFrom:          rule.Validfrom,
				ValidUntil:         rule.Validuntil,
			})
		}
	}

	return &discountRules, nil

}

func (r *DiscountRuleRepository) FindItemsForDiscountRule(ctx context.Context, rule string) ([]string, error) {
	res, err := r.Queries.FindValidItemsForDiscountRule(ctx, rule)
	if err != nil {
		return nil, err
	}

	items := make([]string, 0)
	for _, item := range res {
		items = append(items, item)
	}

	return items, nil
}
