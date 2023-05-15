package rabbit_event

import (
	"github.com/lowl11/lazy-rmq/actors"
	"github.com/lowl11/lazy-rmq/rabbit_service"
)

func (event *Event) ProductionMode() *Event {
	event.isDebug = false
	return event
}

func (event *Event) Publisher() *actors.Publisher {
	return actors.NewPublisher(event.getChannel())
}

func (event *Event) Consumer() *actors.Consumer {
	return actors.NewConsumer(event.getChannel())
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
	connection, channel, err := event.connect()
	if err != nil {
		return err
	}

	event.setConnection(connection, channel)

	for _, consumer := range event.consumers {
		consumer.UpdateChannel(channel)
	}
	return nil
}

func (event *Event) UpdatableConsumer(consumer *actors.Consumer) *Event {
	if consumer == nil {
		return event
	}

	event.consumers = append(event.consumers, consumer)
	return event
}
