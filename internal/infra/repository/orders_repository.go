package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/vemta/common/entity"
	"github.com/vemta/common/enum/orderstatus"
	"github.com/vemta/mvc/internal/infra/db"
)

type OrderRepository struct {
	dbConn *sql.DB
	*db.Queries
	Repository
}

func NewOrderRepository(dbConn *sql.DB) *OrderRepository {
	return &OrderRepository{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

func (r *OrderRepository) Create(ctx context.Context, item *entity.Item) error {
	err := r.Queries.CreateItem(ctx, db.CreateItemParams{
		ID:          item.ID,
		Title:       item.Title,
		Description: item.Description,
		Isgood:      item.IsGood,
		Createdat:   time.Now(),
	})

	return err
}

func (r *OrderRepository) UpdateOrderStatus(ctx context.Context, order string, status int) error {
	return r.Queries.UpdateOrderStatus(ctx, db.UpdateOrderStatusParams{
		ID:     order,
		Status: int32(status),
	})
}

func (r *OrderRepository) FindOrder(ctx context.Context, value string) (*entity.Order, error) {

	orderFound, err := r.Queries.FindOrder(ctx, value)
	if err != nil {
		return nil, err
	}

	order := entity.Order{}
	details := make([]entity.OrderEntry, 0, len(orderFound))

	for _, entry := range orderFound {

		order.Customer = &entity.Customer{
			Email:     entry.Customeremail,
			FullName:  entry.Customerfullname,
			Birthdate: entry.Customerbirthdate,
		}

		order.ID = entry.Orderid
		order.Price = entry.Orderprice
		order.Status = orderstatus.OrderStatus(uint8(entry.Orderstatus))
		order.PaymentMethod = int(entry.Orderpaymentmethod)

		details = append(details, entity.OrderEntry{
			Item: &entity.Item{
				ID:     entry.Itemid,
				Title:  entry.Itemtitle,
				IsGood: entry.Itemisgood,
			},
			Quantity: int64(entry.Detailquantity),
		})
	}

	order.Items = &details

	return &order, nil
}
