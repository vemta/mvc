package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/vemta/common/entity"
	"github.com/vemta/mvc/internal/infra/db"
)

type DiscountRepository struct {
	dbConn *sql.DB
	*db.Queries
	Repository
}

func NewDiscountRepository(dbConn *sql.DB) *DiscountRepository {
	return &DiscountRepository{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

func (r *DiscountRepository) FindItemDiscountRules(ctx context.Context, id string) (*entity.ItemDiscountRule, error) {

	found, err := r.Queries.FindItemDiscountRule(ctx, id)
	if err != nil {
		return nil, err
	}

	applyFirst := ""

	items := make([]string, 0)
	rule := &entity.ItemDiscountRule{}
	rule.ApplyFirst = applyFirst

	for _, entry := range found {
		if err := entry.Applyfirst.Scan(applyFirst); err != nil {
			return nil, err
		}
		validUntil := func() time.Time {
			if entry.Validuntil.Valid {
				return entry.Validuntil.Time
			}
			return time.Time{}
		}
		rule.ID = entry.ID
		rule.DiscountPercentual = entry.Discountpercentual
		rule.DiscountRaw = entry.Discountraw
		rule.AboveValue = entry.Abovevalue
		rule.BellowValue = entry.Bellowvalue
		rule.Name = entry.Name
		rule.ValidUntil = validUntil()
		rule.ValidFrom = entry.Validfrom
		items = append(items, entry.Item)
	}

	rule.Items = items

	return rule, nil
}

func (r *DiscountRepository) FindOrderDiscountRules(ctx context.Context, id string) (*entity.OrderDiscountRule, error) {
	found, err := r.Queries.FindOrderDiscountRule(ctx, id)
	if err != nil {
		return nil, err
	}

	applyFirst := ""
	if err := found.Applyfirst.Scan(applyFirst); err != nil {
		return nil, err
	}

	validUntil := func() time.Time {
		if found.Validuntil.Valid {
			return found.Validuntil.Time
		}
		return time.Time{}
	}

	rule := &entity.OrderDiscountRule{
		ID:                 found.ID,
		ValidFrom:          found.Validfrom,
		ValidUntil:         validUntil(),
		BellowValue:        found.Bellowvalue,
		AboveValue:         found.Abovevalue,
		Name:               found.Name,
		ApplyFirst:         applyFirst,
		DiscountRaw:        found.Discountraw,
		DiscountPercentual: found.Discountpercentual,
	}

	return rule, nil
}

func (r *DiscountRepository) FindValidDiscountRulesForItem(ctx context.Context, id string) (*[]entity.ItemDiscountRule, error) {
	foundRules, err := r.Queries.FindAvailableDiscountRulesForItem(ctx, id)

	if err != nil {
		return nil, err
	}

	rules := make([]entity.ItemDiscountRule, 0)
	for _, rule := range foundRules {

		applyFirst := ""
		if err := rule.Applyfirst.Scan(applyFirst); err != nil {
			return nil, err
		}

		validUntil := func() time.Time {
			if rule.Validuntil.Valid {
				return rule.Validuntil.Time
			}
			return time.Time{}
		}

		rules = append(rules, entity.ItemDiscountRule{
			DiscountPercentual: rule.Discountpercentual,
			DiscountRaw:        rule.Discountraw,
			ID:                 rule.ID,
			ValidFrom:          rule.Validfrom,
			ApplyFirst:         applyFirst,
			Name:               rule.Name,
			AboveValue:         rule.Abovevalue,
			BellowValue:        rule.Bellowvalue,
			ValidUntil:         validUntil(),
		})
	}

	return &rules, nil
}

func (r *DiscountRepository) FindValidOrderDiscountRules(ctx context.Context) (*[]entity.OrderDiscountRule, error) {
	foundRules, err := r.Queries.FindValidOrderDiscountRules(ctx, db.FindValidOrderDiscountRulesParams{
		Validuntil: sql.NullTime{Time: time.Now(), Valid: true},
		Validfrom:  time.Now(),
	})

	if err != nil {
		return nil, err
	}

	rules := make([]entity.ItemDiscountRule, 0)
	for _, rule := range foundRules {

		applyFirst := ""
		if err := rule.Applyfirst.Scan(applyFirst); err != nil {
			return nil, err
		}

		validUntil := func() time.Time {
			if rule.Validuntil.Valid {
				return rule.Validuntil.Time
			}
			return time.Time{}
		}

		rules = append(rules, entity.ItemDiscountRule{
			DiscountPercentual: rule.Discountpercentual,
			DiscountRaw:        rule.Discountraw,
			ID:                 rule.ID,
			ValidFrom:          rule.Validfrom,
			ApplyFirst:         applyFirst,
			Name:               rule.Name,
			AboveValue:         rule.Abovevalue,
			BellowValue:        rule.Bellowvalue,
			ValidUntil:         validUntil(),
		})
	}

	return &rules, nil
}

func (r *DiscountRepository) CreateItemDiscountRule(ctx context.Context, rule entity.ItemDiscountRule) error {
	return errors.New("not implemented yet")
}

func (r *DiscountRepository) CreateOrderDiscountRule(ctx context.Context, rule entity.OrderDiscountRule) error {
	return errors.New("not implemented yet")
}
