package usecase

import (
	"context"

	"github.com/vemta/common/entity"
	"github.com/vemta/mvc/internal/infra/repository"
	uow "github.com/vemta/mvc/pkg"
)

type ItemFindUsecase struct {
	Uow uow.UowInterface
}

func NewItemFindUsecase(uow uow.UowInterface) *ItemFindUsecase {
	return &ItemFindUsecase{
		Uow: uow,
	}
}

func (u *ItemFindUsecase) Execute(ctx context.Context, id string) (*entity.Item, error) {
	return repository.GetItemsRepository(ctx, u.Uow).FindItem(ctx, id)
}
