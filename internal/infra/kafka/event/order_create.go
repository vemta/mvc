package event

import (
	"context"
	"encoding/json"
	"github.com/vemta/mvc/internal/domain/usecase"
	uow "github.com/vemta/mvc/pkg"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type ProcessNewOrder struct{}

func (p ProcessNewOrder) Process(ctx context.Context, msg *kafka.Message, uow uow.UowInterface) error {
	var input usecase.CreateOrderUsecaseInput
	err := json.Unmarshal(msg.Value, &input)
	if err != nil {
		return err
	}
	newOrderUc := usecase.NewCreateOrderUsecase(uow)
	err = newOrderUc.Execute(ctx, input)
	if err != nil {
		return err
	}
	return nil
}
