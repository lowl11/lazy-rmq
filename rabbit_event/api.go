package rabbit_event

import (
	"github.com/lowl11/lazy-rmq/actors"
	"github.com/lowl11/lazy-rmq/rabbit_service"
)

func (event *Event) Publisher() *actors.Publisher {
	if event.publisher != nil {
		return event.publisher
	}

	return actors.NewPublisher(event.channel)
}

func (event *Event) Consumer() *actors.Consumer {
	if event.consumer != nil {
		return event.consumer
	}

	return actors.NewConsumer(event.channel)
}

func (event *Event) Exchange(name, exchangeType string) *rabbit_service.Exchange {
	return rabbit_service.NewExchange(event.channel, name, exchangeType)
}

func (event *Event) Queue(name string) *rabbit_service.Queue {
	return rabbit_service.NewQueue(event.channel, name)
}

func (event *Event) Bind(queue *rabbit_service.Queue, exchange *rabbit_service.Exchange) *rabbit_service.Bind {
	return rabbit_service.NewBind(event.channel, queue, exchange)
}

func (event *Event) Close() error {
	return event.connection.Close()
}
