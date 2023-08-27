package factory

import "github.com/vemta/mvc/internal/infra/kafka/event"

func CreateProcessMessageStrategy(topic string) event.ProcessEventStrategy {
	switch topic {
	case "createPlugin":
		return event.ProcessNewPlugin{}
	case "createUser":
		return event.ProcessNewUser{}
	case "createRelease":
		return event.ProcessNewRelease{}
	}
	return nil
}
