package factory

import "github.com/vemta/mvc/internal/infra/kafka/event"

func CreateProcessMessageStrategy(topic string) event.ProcessEventStrategy {
	switch topic {
	case "create_order":
		return event.ProcessNewOrder{}
	case "create_item":
		return event.ProcessNewItem{}
	case "create_item_discount_rule":
		return event.ProcessNewItemDiscountRule{}
	case "create_order_discount_rule":
		return event.ProcessNewOrderDiscountRule{}
	}
	return nil
}
