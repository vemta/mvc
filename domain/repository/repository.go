package repository

import (
	"context"

	"github.com/vemta/common/entity"
)

type ItemRepositoryInterface interface {
	FindItem(ctx context.Context, id string) (*entity.Item, error)
	Create(ctx context.Context, login *entity.Item) error
}
