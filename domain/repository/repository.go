package repository

import (
	"context"

	"github.com/vemta/common/entity"
)

type ItemRepositoryInterface interface {
	FindItem(ctx context.Context, id string) (*entity.Item, error)
	Create(ctx context.Context, login *entity.Item) error
}

type OrderRepositoryInterface interface {
	FindOrder(ctx context.Context, id string) (*entity.Order, error)
	Create(ctx context.Context, order *entity.Order) error
}
