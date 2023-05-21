package rabbit_event

import (
	"github.com/lowl11/lazy-rmq/rabbit_service"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

func (event *Event) setConnection(connection *amqp.Connection, channel *amqp.Channel) {
	event.mutex.Lock()
	defer event.mutex.Unlock()

	event.connection = connection
	event.channel = channel
}

func (event *Event) closeConnection() error {
	connection := event.getConnection()
	if connection == nil {
		return nil
	}

	return connection.Close()
}

func (event *Event) getConnection() *amqp.Connection {
	event.mutex.Lock()
	defer event.mutex.Unlock()

	if event.connection == nil {
		return nil
	}

	if event.connection.IsClosed() {
		return nil
	}

	return event.connection
}

func (event *Event) getChannel() *amqp.Channel {
	event.mutex.Lock()
	defer event.mutex.Unlock()

	if event.channel == nil || event.channel.IsClosed() {
		if err := event.Reconnect(); err != nil {
			if event.isDebug {
				log.Println("Reconnecting to RabbitMQ error")
			}
		}
	}

	return event.channel
}

func (event *Event) connect() (*amqp.Connection, *amqp.Channel, error) {
	connection, err := rabbit_service.NewConnectionConfig(event.connectionString, time.Second*60)
	if err != nil {
		return nil, nil, err
	}

	channel, err := connection.Channel()
	if err != nil {
		return nil, nil, err
	}

	return connection, channel, nil
}

func (event *Event) connectConfig(heartbeat time.Duration) (*amqp.Connection, *amqp.Channel, error) {
	connection, err := rabbit_service.NewConnectionConfig(event.connectionString, heartbeat)
	if err != nil {
		return nil, nil, err
	}

	channel, err := connection.Channel()
	if err != nil {
		return nil, nil, err
	}

	return connection, channel, nil
}
