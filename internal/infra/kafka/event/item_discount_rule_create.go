package event

import (
	"context"
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/vemta/mvc/internal/domain/usecase"
	uow "github.com/vemta/mvc/pkg"
)

type ProcessNewItemDiscountRule struct{}

func (p ProcessNewItemDiscountRule) Process(ctx context.Context, msg *kafka.Message, uow uow.UowInterface) error {
	var input usecase.CreateItemDiscountRuleUsecaseInput
	err := json.Unmarshal(msg.Value, &input)
	if err != nil {
		return err
	}
	newOrderUc := usecase.NewCreateItemDiscountRuleUsecase(uow)
	err = newOrderUc.Execute(ctx, input)
	if err != nil {
		return err
	}
	return nil
}
