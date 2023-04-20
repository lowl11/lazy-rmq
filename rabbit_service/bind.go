package rabbit_service

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Bind struct {
	channel  *amqp.Channel
	queue    *Queue
	exchange *Exchange

	noWait bool
}

func NewBind(channel *amqp.Channel, queue *Queue, exchange *Exchange) *Bind {
	return &Bind{
		channel:  channel,
		queue:    queue,
		exchange: exchange,
	}
}

func (bind *Bind) NoWait() *Bind {
	bind.noWait = true
	return bind
}

func (bind *Bind) Declare() error {
	if err := bind.channel.QueueBind(
		bind.queue.Name,
		bind.queue.Name,
		bind.exchange.Name,
		bind.noWait,
		nil,
	); err != nil {
		return err
	}
	return nil
}
