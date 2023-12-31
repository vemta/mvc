package repository

import (
	"context"
	"database/sql"

	"github.com/vemta/common/entity"
	"github.com/vemta/common/enum"
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
			currentOrder.PaymentMethod = *enum.GetPaymentMethod(int(order.Paymentmethod))
			currentOrder.Price = order.Orderprice
			currentOrder.Status = enum.OrderStatus(order.Orderstatus)
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
				Category: &entity.ItemCategory{
					ID:   int(order.Itemcategoryid),
					Name: order.Itemcategoryname,
				},
				Valuation: &entity.ItemValuation{
					LastCost:  order.Itemcost,
					LastPrice: order.Itemprice,
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
	found, err := r.Queries.FindCustomer(ctx, customer)
	if err != nil {
		return nil, err
	}

	return &entity.Customer{
		Email:     found.Email,
		FullName:  found.Fullname,
		Birthdate: found.Birthdate,
	}, nil

}

func (r *CustomerRepository) Create(ctx context.Context, customer *entity.Customer) error {
	return r.Queries.CreateCustomer(ctx, db.CreateCustomerParams{
		Email:     customer.Email,
		Fullname:  customer.FullName,
		Birthdate: customer.Birthdate,
	})
}
