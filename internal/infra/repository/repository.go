package repository

import (
	"context"
	"errors"

	"github.com/vemta/mvc/domain/repository"
	"github.com/vemta/mvc/internal/infra/db"
	uow "github.com/vemta/mvc/pkg"
)

var ErrQueriesNotSet = errors.New("queries not set")

type Repository struct {
	*db.Queries
}

func (r *Repository) SetQuery(q *db.Queries) {
	r.Queries = q
}

func (r *Repository) Validate() error {
	if r.Queries == nil {
		return ErrQueriesNotSet
	}
	return nil
}

func GetItemsRepository(ctx context.Context, u uow.UowInterface) repository.ItemsRepositoryInterface {
	return getRepository[repository.ItemsRepositoryInterface](ctx, u, "ItemsRepository")
}

func GetCustomersRepository(ctx context.Context, u uow.UowInterface) repository.CustomersRepositoryInterface {
	return getRepository[repository.CustomersRepositoryInterface](ctx, u, "CustomersRepository")
}

func GetOrdersRepository(ctx context.Context, u uow.UowInterface) repository.OrdersRepositoryInterface {
	return getRepository[repository.OrdersRepositoryInterface](ctx, u, "OrdersRepository")
}

func getRepository[T repository.RepositoryInterface](ctx context.Context, u uow.UowInterface, name string) T {
	repo, err := u.GetRepository(ctx, name)
	if err != nil {
		panic(err)
	}
	return repo.(T)
}
