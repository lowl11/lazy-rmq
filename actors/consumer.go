package actors

import (
	"github.com/lowl11/lazy-rmq/rabbit_service"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

type Consumer struct {
	channel *amqp.Channel

	errorHandler func(err error)
}

func NewConsumer(channel *amqp.Channel) *Consumer {
	return &Consumer{
		channel: channel,
	}
}

func (consumer *Consumer) ErrorHandler(handler func(err error)) *Consumer {
	consumer.errorHandler = handler
	return consumer
}

func (consumer *Consumer) Handler(queue *rabbit_service.Queue, handler rabbit_service.HandlerFunc) error {
	messages, err := rabbit_service.
		NewHandler(consumer.channel, queue).
		Declare()
	if err != nil {
		return err
	}

	go consumer.runHandler(messages, handler)
	return nil
}

func (consumer *Consumer) runHandler(messages <-chan amqp.Delivery, handler rabbit_service.HandlerFunc) {
	for message := range messages {
		handlerContext := &rabbit_service.HandlerContext{
			Delivery: message,
		}
		if err := handler(handlerContext); err != nil {
			if consumer.errorHandler != nil {
				consumer.errorHandler(err)
			} else {
				log.Println("Error from message handler: ", err)
			}
		}
	}
}
