package event

import (
	"context"
	uow "github.com/vemta/mvc/pkg"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type ProcessEventStrategy interface {
	Process(ctx context.Context, msg *kafka.Message, uow uow.UowInterface) error
}
