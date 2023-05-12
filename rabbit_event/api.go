package rabbit_event

import (
	"github.com/lowl11/lazy-rmq/actors"
	"github.com/lowl11/lazy-rmq/rabbit_service"
)

func (event *Event) Publisher() *actors.Publisher {
	if event.publisher != nil {
		return event.publisher
	}

	event.publisher = actors.NewPublisher(event.getChannel())
	return event.publisher
}

func (event *Event) Consumer() *actors.Consumer {
	if event.consumer != nil {
		return event.consumer
	}

	event.consumer = actors.NewConsumer(event.getChannel())
	return event.consumer
}

func (event *Event) Exchange(name, exchangeType string) *rabbit_service.Exchange {
	return rabbit_service.NewExchange(event.getChannel(), name, exchangeType)
}

func (event *Event) Queue(name string) *rabbit_service.Queue {
	return rabbit_service.NewQueue(event.getChannel(), name)
}

func (event *Event) Bind(queue *rabbit_service.Queue, exchange *rabbit_service.Exchange) *rabbit_service.Bind {
	return rabbit_service.NewBind(event.getChannel(), queue, exchange)
}

func (event *Event) Close() error {
	return event.closeConnection()
}

func (event *Event) IsClosed() bool {
	connection := event.getConnection()
	if connection == nil {
		return true
	}

	return connection.IsClosed()
}

func (event *Event) Reconnect() error {
	connection, err := rabbit_service.NewConnection(event.connectionString)
	if err != nil {
		return err
	}

	channel, err := connection.Channel()
	if err != nil {
		return err
	}

	event.setConnection(connection, channel)
	event.publisher = nil
	event.consumer = nil
	return nil
}
