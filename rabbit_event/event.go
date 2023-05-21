package rabbit_event

import (
	"github.com/lowl11/lazy-rmq/actors"
	amqp "github.com/rabbitmq/amqp091-go"
	"sync"
	"time"
)

type Event struct {
	connectionString string

	channel    *amqp.Channel
	connection *amqp.Connection

	heartbeat time.Duration
	isDebug   bool

	consumers []*actors.Consumer

	mutex sync.Mutex
}

func New(connectionString string) (*Event, error) {
	event := &Event{
		connectionString: connectionString,
		consumers:        make([]*actors.Consumer, 0),
		heartbeat:        time.Second * 60,
		isDebug:          true,
	}

	connection, channel, err := event.connect()
	if err != nil {
		return nil, err
	}

	event.setConnection(connection, channel)

	return event, nil
}

func NewConfig(connectionString string, heartbeat time.Duration) (*Event, error) {
	event := &Event{
		connectionString: connectionString,
		consumers:        make([]*actors.Consumer, 0),
		heartbeat:        heartbeat,
		isDebug:          true,
	}

	connection, channel, err := event.connectConfig(heartbeat)
	if err != nil {
		return nil, err
	}

	event.setConnection(connection, channel)

	return event, nil
}
