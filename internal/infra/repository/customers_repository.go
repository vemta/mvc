package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/vemta/common/entity"
	"github.com/vemta/mvc/internal/infra/db"
)

type CustomerRepository struct {
	dbConn *sql.DB
	*db.Queries
	Repository
}

func NewCustomerRepository(dbConn *sql.DB) *CustomerRepository {
	return &CustomerRepository{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

func (r *CustomerRepository) FindCustomerOrders(ctx context.Context, customer string) (*[]entity.Order, error) {
	orders, err := r.Queries.FindCustomerOrders(ctx, customer)
	if err != nil {
		return nil, err
	}

	ordersSlice := make([]entity.Order, 0)
	currentOrder := &entity.Order{
		ID: "",
	}

	items := make([]entity.OrderEntry, 0)
	for _, order := range orders {
		if currentOrder.ID != order.Orderid {
			if currentOrder.ID != "" {
				itemsCopy := make([]entity.OrderEntry, 0)
				copy(itemsCopy, items)
				currentOrder.Items = &itemsCopy
				items = make([]entity.OrderEntry, 0)
				ordersSlice = append(ordersSlice, *currentOrder)
			}
			currentOrder.ID = order.Orderid
			currentOrder.DiscountRaw = order.Discountraw
			currentOrder.DiscountPercentual = order.Discountpercentual
			currentOrder.PaymentMethod = int(order.Paymentmethod)
			currentOrder.Price = order.Orderprice
			currentOrder.Status = uint8(order.Orderstatus)
			currentOrder.Customer = &entity.Customer{
				Email:     order.Customeremail,
				FullName:  order.Customerfullname,
				Birthdate: order.Customerbirthdate,
			}
		}

		items = append(items, entity.OrderEntry{
			Item: &entity.Item{
				ID:          order.Itemid,
				Title:       order.Itemtitle,
				IsGood:      order.Itemisgood,
				Description: order.Itemdescription,
				Valuation: &entity.ItemValuation{
					DiscountRaw:        order.Itemdiscountraw,
					LastCost:           order.Itemcost,
					LastPrice:          order.Itemprice,
					DiscountPercentual: order.Itemdiscountpercentual,
				},
			},
		})
	}
	itemsCopy := make([]entity.OrderEntry, 0)
	copy(itemsCopy, items)
	currentOrder.Items = &itemsCopy
	ordersSlice = append(ordersSlice, *currentOrder)

	return &ordersSlice, nil
}

func (r *CustomerRepository) FindCustomer(ctx context.Context, customer string) (*entity.Customer, error) {
	return nil, errors.New("not implemented yet")
}

func (r *CustomerRepository) Create(ctx context.Context, customer *entity.Customer) error {
	return errors.New("not implemented yet")
}
