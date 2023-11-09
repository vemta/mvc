package repository

import (
	"context"
	"time"

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

type DiscountRuleRepositoryInterface interface {
	FindValidDiscountRulesForItem(context.Context, string, time.Time) (*[]entity.ItemDiscountRule, error)
	FindValidDiscountRulesForOrder(context.Context, float64, time.Time) (*[]entity.OrderDiscountRule, error)
	FindActiveDiscountRules(context.Context, time.Time) (*[]any, error)
	FindItemsForDiscountRule(context.Context, string) ([]string, error)
	FindAppliedDiscountsForOrder(context.Context, string) (*[]entity.DiscountRule, error)
	FindAutoApplyDiscountRulesForItem(context.Context, string, time.Time) (*[]entity.DiscountRule, error)
	FindAutoApplyDiscountRulesForOrder(context.Context, float64, time.Time) (*[]entity.DiscountRule, error)
	CreateItemDiscountRule(context.Context, *entity.ItemDiscountRule) error
	CreateOrderDiscountRule(context.Context, *entity.OrderDiscountRule) error
}
