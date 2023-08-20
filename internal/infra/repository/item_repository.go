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

	return &entity.Item{
		ID:          item.ID,
		Title:       item.Title,
		IsGood:      item.Isgood,
		Description: item.Description,
	}, nil
}
