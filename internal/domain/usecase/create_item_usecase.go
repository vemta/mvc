package usecase

import (
	"context"

	"github.com/vemta/common/entity"
	"github.com/vemta/mvc/internal/infra/repository"
	uow "github.com/vemta/mvc/pkg"
)

type CreateItemUsecaseInput struct {
	ID          string                `json:"id"`
	Title       string                `json:"title"`
	IsGood      bool                  `json:"is_good"`
	Description string                `json:"description"`
	Category    *entity.ItemCategory  `json:"category"`
	Valuation   *entity.ItemValuation `json:"valuation"`
}

type CreateItemUsecase struct {
	Uow uow.UowInterface
}

func NewCreateItemUsecase(uow uow.UowInterface) *CreateItemUsecase {
	return &CreateItemUsecase{
		Uow: uow,
	}
}

func (u *CreateItemUsecase) Execute(ctx context.Context, input CreateItemUsecaseInput) error {

	return repository.GetItemsRepository(ctx, uow.GetCurrent()).Create(ctx, &entity.Item{
		ID:          input.ID,
		Category:    input.Category,
		IsGood:      input.IsGood,
		Title:       input.Title,
		Description: input.Description,
		Valuation:   input.Valuation,
	})

}
