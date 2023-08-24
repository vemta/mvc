package repository

import (
	"context"
	"database/sql"

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

	currentOrder := &entity.Order{
		ID: "",
	}

	items := make([]entity.OrderEntry, 0)
	for _, order := range orders {
		if currentOrder.ID != order.Orderid {
			if currentOrder.ID == "" {
				currentOrder.ID = order.Orderid
				currentOrder.Customer = &entity.Customer{
					Email: 
				}
			} else {

			}
		}
		items = append(items, order)
	}

}

func (r *CustomerRepository) FindCustomer(ctx context.Context, customer string) (*entity.Customer, error) {

}

func (r *CustomerRepository) Create(ctx context.Context, customer *entity.Customer) error {

}
