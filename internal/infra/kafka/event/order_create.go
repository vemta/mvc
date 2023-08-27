package event

import (
	"context"
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	usecase "github.com/vemta/mvc/internal/domain/usecase/order"
)

type ProcessNewPlugin struct{}

func (p ProcessNewPlugin) Process(ctx context.Context, msg *kafka.Message, uow uow.UowInterface) error {
	var input usecase.
	err := json.Unmarshal(msg.Value, &input)
	if err != nil {
		return err
	}
	addNewPlugin := usecase.NewPluginAddUseCase(uow)
	err = addNewPlugin.Execute(ctx, input)
	if err != nil {
		return err
	}
	return nil
}
