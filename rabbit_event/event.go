package rabbit_event

import (
	"github.com/lowl11/lazy-rmq/actors"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Event struct {
	connection *amqp.Connection
	channel    *amqp.Channel

	publisher *actors.Publisher
	consumer  *actors.Consumer
}

func New(connection *amqp.Connection) (*Event, error) {
	channel, err := connection.Channel()
	if err != nil {
		return nil, err
	}

	return &Event{
		connection: connection,
		channel:    channel,
	}, nil
}
