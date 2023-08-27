package factory

import "github.com/vemta/mvc/internal/infra/kafka/event"

func CreateProcessMessageStrategy(topic string) event.ProcessEventStrategy {
	switch topic {
	case "create_order":
		return event.ProcessNewOrder{}
	}
	return nil
}
