package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/vemta/common/entity"
	"github.com/vemta/mvc/internal/infra/db"
)

type ItemRepository struct {
	dbConn *sql.DB
	*db.Queries
	Repository
}

func NewItemRepository(dbConn *sql.DB) *ItemRepository {
	return &ItemRepository{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

func (r *ItemRepository) Create(ctx context.Context, item *entity.Item) error {
	err := r.Queries.CreateItem(ctx, db.CreateItemParams{
		ID:          item.ID,
		Title:       item.Title,
		Description: item.Description,
		Isgood:      item.IsGood,
		Createdat:   time.Now(),
	})

	return err
}

func (r *ItemRepository) FindItem(ctx context.Context, value string) (*entity.Item, error) {
	item, err := r.Queries.FindItem(ctx, value)

	if err != nil {
		return nil, err
	}

	costHistory, err := r.FindItemCostHistory(ctx, value)
	if err != nil {
		return nil, err
	}

	priceHistory, err := r.FindItemPriceHistory(ctx, value)
	if err != nil {
		return nil, err
	}

	return &entity.Item{
		ID:          item.ID,
		Title:       item.Title,
		IsGood:      item.Isgood,
		Description: item.Description,
		Valuation: &entity.ItemValuation{
			PriceHistory:       *priceHistory,
			CostHistory:        *costHistory,
			DiscountRaw:        item.Discountraw,
			DiscountPercentual: item.Discountpercentual,
			LastCost:           item.Lastcost,
			LastPrice:          item.Lastprice,
		},
	}, nil
}

func (r *ItemRepository) FindItemCostHistory(ctx context.Context, value string) (*[]entity.ItemValuationLog, error) {
	logs, err := r.Queries.FindItemCostHistory(ctx, value)

	if err != nil {
		return nil, err
	}

	log := make([]entity.ItemValuationLog, 0, len(logs))

	for _, entry := range logs {
		log = append(log, entity.ItemValuationLog{
			Value:              entry.Price,
			DiscountRaw:        entry.Discountraw,
			DiscountPercentual: entry.Discountpercentual,
			UpdatedAt:          entry.Valorizatedat,
		})
	}

	return &log, nil
}

func (r *ItemRepository) FindItemPriceHistory(ctx context.Context, value string) (*[]entity.ItemValuationLog, error) {
	logs, err := r.Queries.FindItemPriceHistory(ctx, value)

	if err != nil {
		return nil, err
	}

	log := make([]entity.ItemValuationLog, 0, len(logs))

	for _, entry := range logs {
		log = append(log, entity.ItemValuationLog{
			Value:              entry.Price,
			DiscountRaw:        entry.Discountraw,
			DiscountPercentual: entry.Discountpercentual,
			UpdatedAt:          entry.Valorizatedat,
		})
	}

	return &log, nil
}

func (r *ItemRepository) UpdateItemValorization(ctx context.Context, item string, valuation *entity.ItemValuation) error {
	return r.Queries.UpdateItemValorization(ctx, db.UpdateItemValorizationParams{
		Lastprice:          valuation.LastPrice,
		Lastcost:           valuation.LastCost,
		Discountraw:        valuation.DiscountRaw,
		Discountpercentual: valuation.DiscountPercentual,
		Updatedat:          valuation.UpdatedAt,
		Itemid:             item,
	})
}
