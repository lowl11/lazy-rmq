package rabbit_event

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func (event *Event) setConnection(connection *amqp.Connection, channel *amqp.Channel) {
	event.mutex.Lock()
	defer event.mutex.Unlock()

	event.connection = connection
	event.channel = channel
}

func (event *Event) closeConnection() error {
	event.mutex.Lock()
	defer event.mutex.Unlock()

	if event.connection == nil {
		return nil
	}

	return event.connection.Close()
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
	return event.channel
}
