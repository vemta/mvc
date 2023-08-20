package usecase

import (
	"context"

	"github.com/vemta/common/entity"
	"github.com/vemta/mvc/domain/repository"
	uow "github.com/vemta/mvc/pkg"
)

type ItemFindUsecaseInput struct {
	ID string `json:"id"`
}

type ItemFindUsecase struct {
	Uow uow.UowInterface
}

func NewItemFindUsecase(uow uow.UowInterface) *ItemFindUsecase {
	return &ItemFindUsecase{
		Uow: uow,
	}
}

func (u *ItemFindUsecase) Execute(ctx context.Context, input ItemFindUsecaseInput) (*entity.Item, error) {
	return u.getItemRepository(ctx).FindItem(ctx, input.ID)
}

func (u *ItemFindUsecase) getItemRepository(ctx context.Context) repository.ItemRepositoryInterface {
	itemRepository, err := u.Uow.GetRepository(ctx, "ItemRepository")
	if err != nil {
		panic(err)
	}
	return itemRepository.(repository.ItemRepositoryInterface)
}
