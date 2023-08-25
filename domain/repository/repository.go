package repository

import (
	"context"

	"github.com/vemta/common/entity"
)

type RepositoryInterface interface {
}

type ItemsRepositoryInterface interface {
	FindItem(context.Context, string) (*entity.Item, error)
	FindItemCostHistory(context.Context, string) (*[]entity.ItemValuationLog, error)
	FindItemPriceHistory(context.Context, string) (*[]entity.ItemValuationLog, error)
	UpdateItemValuation(context.Context, string, *entity.ItemValuation) error
	Create(ctx context.Context, login *entity.Item) error
	RepositoryInterface
}

type OrdersRepositoryInterface interface {
	FindOrder(context.Context, string) (*entity.Order, error)
	Create(context.Context, *entity.Order) error
	RepositoryInterface
}

type CustomersRepositoryInterface interface {
	FindCustomerOrders(context.Context, string) (*[]entity.Order, error)
	FindCustomer(context.Context, string) (*entity.Customer, error)
	Create(context.Context, *entity.Customer) error
	RepositoryInterface
}
