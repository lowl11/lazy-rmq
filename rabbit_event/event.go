package rabbit_event

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"sync"
)

type Event struct {
	connectionString string

	channel    *amqp.Channel
	connection *amqp.Connection

	isDebug bool

	mutex sync.Mutex
}

func New(connectionString string) (*Event, error) {
	event := &Event{
		connectionString: connectionString,
		isDebug:          true,
	}

	connection, channel, err := event.connect()
	if err != nil {
		return nil, err
	}

	event.setConnection(connection, channel)

	return event, nil
}
