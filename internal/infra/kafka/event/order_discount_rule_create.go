package event

import (
	"context"
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/vemta/mvc/internal/domain/usecase"
	uow "github.com/vemta/mvc/pkg"
)

type ProcessNewOrderDiscountRule struct{}

func (p ProcessNewOrderDiscountRule) Process(ctx context.Context, msg *kafka.Message, uow uow.UowInterface) error {
	var input usecase.CreateOrderDiscountRuleUsecaseInput
	err := json.Unmarshal(msg.Value, &input)
	if err != nil {
		return err
	}
	newOrderUc := usecase.NewCreateOrderDiscountRuleUsecase(uow)
	err = newOrderUc.Execute(ctx, input)
	if err != nil {
		return err
	}
	return nil
}
