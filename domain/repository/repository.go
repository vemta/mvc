package repository

import (
	"context"

	"github.com/vemta/common/entity"
)

type RepositoryInterface interface {
}

type ItemsRepositoryInterface interface {
	FindItem(ctx context.Context, id string) (*entity.Item, error)
	Create(ctx context.Context, login *entity.Item) error
	RepositoryInterface
}

type OrdersRepositoryInterface interface {
	FindOrder(ctx context.Context, id string) (*entity.Order, error)
	Create(ctx context.Context, order *entity.Order) error
	RepositoryInterface
}

type CustomersRepositoryInterface interface {
	FindCustomerOrders(ctx context.Context, customer string) (*[]entity.Order, error)
	FindCustomer(ctx context.Context, customer string) (*entity.Customer, error)
	Create(ctx context.Context, customer *entity.Customer) error
	RepositoryInterface
}
