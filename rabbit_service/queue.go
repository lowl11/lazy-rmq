package rabbit_service

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Queue struct {
	Name string

	channel    *amqp.Channel
	createName string
	durable    bool
	autoDelete bool
	exclusive  bool
	noWait     bool
}

func NewQueue(channel *amqp.Channel, name string) *Queue {
	return &Queue{
		channel:    channel,
		createName: name,
	}
}

func (queue *Queue) Durable() *Queue {
	queue.durable = true
	return queue
}

func (queue *Queue) AutoDelete() *Queue {
	queue.autoDelete = true
	return queue
}

func (queue *Queue) Exclusive() *Queue {
	queue.exclusive = true
	return queue
}

func (queue *Queue) NoWait() *Queue {
	queue.noWait = true
	return queue
}

func (queue *Queue) Declare() (*Queue, error) {
	createdQueue, err := queue.channel.QueueDeclare(
		queue.createName,
		queue.durable,
		queue.autoDelete,
		queue.exclusive,
		queue.noWait,
		nil,
	)
	if err != nil {
		return queue, err
	}

	queue.Name = createdQueue.Name
	return queue, nil
}
