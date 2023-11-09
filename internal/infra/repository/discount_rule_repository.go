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
			DiscountRule: &entity.DiscountRule{
				ID:                 rule.ID,
				Name:               rule.Name,
				DiscountRaw:        rule.Discountraw,
				DiscountPercentual: rule.Discountpercentual,
				ApplyFirst:         rule.Applyfirst,
				AboveValue:         rule.Abovevalue,
				BellowValue:        rule.Abovevalue,
				ValidFrom:          rule.Validfrom,
				ValidUntil:         rule.Validuntil,
			},
			Items: []string{rule.Item},
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
		code := ""
		if rule.Code.Valid {
			code = rule.Code.String
		}
		discountRules = append(discountRules, entity.OrderDiscountRule{
			DiscountRule: &entity.DiscountRule{
				ID:                 rule.ID,
				Name:               rule.Name,
				DiscountRaw:        rule.Discountraw,
				DiscountPercentual: rule.Discountpercentual,
				ApplyFirst:         rule.Applyfirst,
				AboveValue:         rule.Abovevalue,
				BellowValue:        rule.Abovevalue,
				ValidFrom:          rule.Validfrom,
				ValidUntil:         rule.Validuntil,
				Code:               code,
				AutoApply:          rule.Autoapply != 0,
			},
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

func (r *DiscountRuleRepository) FindActiveDiscountRules(ctx context.Context, time time.Time) (*[]any, error) {

	rules, err := r.Queries.FindActiveDiscountRules(ctx, db.FindActiveDiscountRulesParams{
		Validfrom:  time,
		Validuntil: time,
	})

	if err != nil {
		return nil, err
	}

	discountRules := make([]any, 0)

	for _, rule := range rules {

		discountRule := &entity.DiscountRule{
			ID:                 rule.ID,
			Name:               rule.Name,
			DiscountRaw:        rule.Discountraw,
			DiscountPercentual: rule.Discountpercentual,
			ApplyFirst:         rule.Applyfirst,
			AboveValue:         rule.Abovevalue,
			BellowValue:        rule.Bellowvalue,
			ValidFrom:          rule.Validfrom,
			ValidUntil:         rule.Validuntil,
		}

		switch []rune(rule.Type)[0] {
		case 'I':
			items, err := r.FindItemsForDiscountRule(ctx, rule.ID)
			if err != nil {
				return nil, err
			}
			discountRules = append(discountRules, entity.ItemDiscountRule{
				DiscountRule: discountRule,
				Items:        items,
			})
		case 'O':
			discountRules = append(discountRules, entity.OrderDiscountRule{
				DiscountRule: discountRule,
			})
		}
	}
	return &discountRules, nil
}

func (r *DiscountRuleRepository) FindItemsForDiscountRule(ctx context.Context, rule string) ([]string, error) {
	res, err := r.Queries.FindValidItemsForDiscountRuleDetailed(ctx, rule)
	if err != nil {
		return nil, err
	}

	items := make([]string, 0)
	for _, item := range res {
		items = append(items, item.Itemid)
	}

	return items, nil
}

func (r *DiscountRuleRepository) FindAppliedDiscountsForOrder(ctx context.Context, order string) (*[]entity.DiscountRule, error) {

	discounts, err := r.Queries.FindAppliedDiscountRulesForOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	rules := make([]entity.DiscountRule, 0)

	for _, discount := range discounts {
		code := ""
		if discount.Discountcode.Valid {
			code = discount.Discountcode.String
		}
		rules = append(rules, entity.DiscountRule{
			ID:                 discount.Discountrule,
			Code:               code,
			AutoApply:          discount.Autoapply == 1,
			Name:               discount.Discountname,
			DiscountRaw:        discount.Discountraw,
			DiscountPercentual: discount.Discountpercentual,
			ApplyFirst:         discount.Applyfirst,
			AboveValue:         discount.Abovevalue,
			BellowValue:        discount.Bellowvalue,
			ValidFrom:          discount.Validfrom,
			ValidUntil:         discount.Validuntil,
			Type:               []rune(discount.Discounttype)[0],
		})
	}

	return &rules, nil
}

func (r *DiscountRuleRepository) FindAutoApplyDiscountRulesForItem(ctx context.Context, item string, time time.Time) (*[]entity.DiscountRule, error) {
	discounts, err := r.Queries.FindAutoApplyDiscountRulesForItem(ctx, db.FindAutoApplyDiscountRulesForItemParams{
		Item:       item,
		Validfrom:  time,
		Validuntil: time,
	})
	if err != nil {
		return nil, err
	}

	rules := make([]entity.DiscountRule, 0)

	for _, discount := range discounts {
		code := ""
		if discount.Discountcode.Valid {
			code = discount.Discountcode.String
		}
		rules = append(rules, entity.DiscountRule{
			ID:                 discount.Discountrule,
			Code:               code,
			AutoApply:          discount.Autoapply == 1,
			Name:               discount.Discountname,
			DiscountRaw:        discount.Discountraw,
			DiscountPercentual: discount.Discountpercentual,
			ApplyFirst:         discount.Applyfirst,
			AboveValue:         discount.Abovevalue,
			BellowValue:        discount.Bellowvalue,
			ValidFrom:          discount.Validfrom,
			ValidUntil:         discount.Validuntil,
			Type:               []rune(discount.Discounttype)[0],
		})
	}

	return &rules, nil
}

func (r *DiscountRuleRepository) FindAutoApplyDiscountRulesForOrder(ctx context.Context, price float64, time time.Time) (*[]entity.DiscountRule, error) {
	discounts, err := r.Queries.FindAutoApplyDiscountRulesForOrder(ctx, db.FindAutoApplyDiscountRulesForOrderParams{
		Validfrom:   time,
		Validuntil:  time,
		Bellowvalue: price,
		Abovevalue:  price,
	})
	if err != nil {
		return nil, err
	}

	rules := make([]entity.DiscountRule, 0)

	for _, discount := range discounts {
		code := ""
		if discount.Discountcode.Valid {
			code = discount.Discountcode.String
		}
		rules = append(rules, entity.DiscountRule{
			ID:                 discount.Discountrule,
			Code:               code,
			AutoApply:          discount.Autoapply == 1,
			Name:               discount.Discountname,
			DiscountRaw:        discount.Discountraw,
			DiscountPercentual: discount.Discountpercentual,
			ApplyFirst:         discount.Applyfirst,
			AboveValue:         discount.Abovevalue,
			BellowValue:        discount.Bellowvalue,
			ValidFrom:          discount.Validfrom,
			ValidUntil:         discount.Validuntil,
			Type:               []rune(discount.Discounttype)[0],
		})
	}

	return &rules, nil
}
