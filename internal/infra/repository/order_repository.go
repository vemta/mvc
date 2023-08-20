package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/vemta/common/entity"
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

func (r *OrderRepository) FindOrder(ctx context.Context, value string) (*entity.Order, error) {

	orderFound, err := r.Queries.FindOrder(ctx, value)
	if err != nil {
		return nil, err
	}

	order := entity.Order{}
	details := make([]entity.OrderDetail, 0, len(orderFound))

	for _, entry := range orderFound {

		order.Customer = &entity.User{
			Email:     entry.Customeremail,
			FullName:  entry.Customerfullname,
			Birthdate: entry.Customerbirthdate,
		}

		order.DiscountPercentual = entry.Orderdiscountpercentual
		order.ID = entry.Orderid
		order.DiscountRaw = entry.Orderdiscountraw
		order.Price = entry.Orderprice
		order.Status = uint8(entry.Orderstatus)
		order.PaymentMethod = int(entry.Orderpaymentmethod)

		details = append(details, entity.OrderDetail{
			Item: &entity.Item{
				ID:     entry.Itemid,
				Title:  entry.Itemtitle,
				IsGood: entry.Itemisgood,
			},
			Quantity: int(entry.Detailquantity.Int32),
		})
	}

	if err != nil {
		return nil, err
	}

	return &entity.Item{
		ID:          item.ID,
		Title:       item.Title,
		IsGood:      item.Isgood,
		Description: item.Description,
	}, nil
}
