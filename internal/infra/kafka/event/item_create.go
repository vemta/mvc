package event

import (
	"context"
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/vemta/mvc/internal/domain/usecase"
	uow "github.com/vemta/mvc/pkg"
)

type ProcessNewItem struct{}

func (p ProcessNewItem) Process(ctx context.Context, msg *kafka.Message, uow uow.UowInterface) error {
	var input usecase.CreateItemUsecaseInput
	err := json.Unmarshal(msg.Value, &input)
	if err != nil {
		return err
	}
	newOrderUc := usecase.NewCreateItemUsecase(uow)
	err = newOrderUc.Execute(ctx, input)
	if err != nil {
		return err
	}
	return nil
}
