package rabbit_service

import amqp "github.com/rabbitmq/amqp091-go"

func NewConnection(connectionString string) (*amqp.Connection, error) {
	connection, err := amqp.Dial(connectionString)
	if err != nil {
		return nil, err
	}
	return connection, nil
}
