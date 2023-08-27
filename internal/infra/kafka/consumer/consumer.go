package consumer

import (
	"context"
	"fmt"
	uow "github.com/vemta/mvc/pkg"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/vemta/mvc/internal/infra/kafka/factory"
)

func Consume(topics []string, servers string, msgChan chan *kafka.Message) {
	kafkaConsumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": servers,
		"group.id":          "gostats",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}
	kafkaConsumer.SubscribeTopics(topics, nil)
	for {
		msg, err := kafkaConsumer.ReadMessage(-1)
		if err == nil {
			msgChan <- msg
		}
	}
}

func ProcessEvents(ctx context.Context, msgChan chan *kafka.Message, uwo uow.UowInterface) {
	for msg := range msgChan {
		fmt.Println("Received message: ", string(msg.Value), " in topic ", msg.TopicPartition.Topic)
		strategy := factory.CreateProcessMessageStrategy(*msg.TopicPartition.Topic)
		err := strategy.Process(ctx, msg, uwo)
		if err != nil {
			fmt.Println(err)
		}
	}
}
