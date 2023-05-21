package rabbit_service

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

func NewConnection(connectionString string) (*amqp.Connection, error) {
	connection, err := amqp.Dial(connectionString)
	if err != nil {
		return nil, err
	}

	return connection, nil
}

func NewConnectionConfig(connectionString string, heartbeat time.Duration) (*amqp.Connection, error) {
	connection, err := amqp.DialConfig(connectionString, amqp.Config{
		Heartbeat: heartbeat,
	})
	if err != nil {
		return nil, err
	}

	return connection, nil
}
