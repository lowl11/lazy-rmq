package rabbit_service

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type HandlerFunc func(ctx *HandlerContext) error

type HandlerContext struct {
	amqp.Delivery
}

type Handler struct {
	channel *amqp.Channel
	queue   *Queue

	consumer  string
	autoAck   bool
	exclusive bool
	noLocal   bool
	noWait    bool
}

func NewHandler(channel *amqp.Channel, queue *Queue) *Handler {
	return &Handler{
		channel: channel,
		queue:   queue,
		autoAck: true,
	}
}

func (handler *Handler) Consumer(consumer string) *Handler {
	handler.consumer = consumer
	return handler
}

func (handler *Handler) AutoAck() *Handler {
	handler.autoAck = true
	return handler
}

func (handler *Handler) Exclusive() *Handler {
	handler.exclusive = true
	return handler
}

func (handler *Handler) NoLocal() *Handler {
	handler.noLocal = true
	return handler
}

func (handler *Handler) NoWait() *Handler {
	handler.noWait = true
	return handler
}

func (handler *Handler) Declare() (<-chan amqp.Delivery, error) {
	messagesChan, err := handler.channel.Consume(
		handler.queue.Name,
		handler.consumer,
		handler.autoAck,
		handler.exclusive,
		handler.noLocal,
		handler.noWait,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return messagesChan, nil
}
