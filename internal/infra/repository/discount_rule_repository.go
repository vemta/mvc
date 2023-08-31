package repository

import (
	"context"
	"database/sql"
	"errors"

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
	return nil, errors.New("not implemented yet")
}

func (r *DiscountRepository) FindOrderDiscountRules(ctx context.Context, id string) (*entity.OrderDiscountRule, error) {
	return nil, errors.New("not implemented yet")
}

func (r *DiscountRepository) FindValidDiscountRulesForItem(ctx context.Context, id string) (*[]entity.ItemDiscountRule, error) {
	return nil, errors.New("not implemented yet")
}

func (r *DiscountRepository) FindValidOrderDiscountRules(ctx context.Context) (*[]entity.OrderDiscountRule, error) {
	return nil, errors.New("not implemented yet")
}

func (r *DiscountRepository) CreateItemDiscountRule(ctx context.Context, rule entity.ItemDiscountRule) error {
	return errors.New("not implemented yet")
}

func (r *DiscountRepository) CreateOrderDiscountRule(ctx context.Context, rule entity.OrderDiscountRule) error {
	return errors.New("not implemented yet")
}
