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
	Create(context.Context, *entity.Item) error
	RepositoryInterface
}

type OrdersRepositoryInterface interface {
	FindOrder(context.Context, string) (*entity.Order, error)
	FindOrderFinalPrice(context.Context, string) (float64, error)
	UpdateOrderStatus(context.Context, string, int) error
	Create(context.Context, *entity.Order) error
	RepositoryInterface
}

type CustomersRepositoryInterface interface {
	FindCustomerOrders(context.Context, string) (*[]entity.Order, error)
	FindCustomer(context.Context, string) (*entity.Customer, error)
	Create(context.Context, *entity.Customer) error
	RepositoryInterface
}

type DiscountRulesRepositoryInterface interface {
	FindItemDiscountRules(context.Context, string) (*entity.ItemDiscountRule, error)
	FindOrderDiscountRules(context.Context, string) (*entity.OrderDiscountRule, error)
	FindValidDiscountRulesForItem(context.Context, string) (*[]entity.ItemDiscountRule, error)
	FindValidOrderDiscountRules(context.Context) (*[]entity.OrderDiscountRule, error)
	CreateItemDiscountRule(context.Context, *entity.ItemDiscountRule) error
	CreateOrderDiscountRule(context.Context, *entity.OrderDiscountRule) error
}
