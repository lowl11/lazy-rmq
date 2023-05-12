package rabbit_event

import (
	"github.com/lowl11/lazy-rmq/actors"
	"github.com/lowl11/lazy-rmq/rabbit_service"
	amqp "github.com/rabbitmq/amqp091-go"
	"sync"
)

type Event struct {
	connectionString string

	connection *amqp.Connection
	channel    *amqp.Channel

	publisher *actors.Publisher
	consumer  *actors.Consumer

	mutex sync.Mutex
}

func New(connectionString string) (*Event, error) {
	connection, err := rabbit_service.NewConnection(connectionString)
	if err != nil {
		return nil, err
	}

	channel, err := connection.Channel()
	if err != nil {
		return nil, err
	}

	return &Event{
		connectionString: connectionString,
		connection:       connection,
		channel:          channel,
	}, nil
}
