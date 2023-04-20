package rabbit_service

import amqp "github.com/rabbitmq/amqp091-go"

type Exchange struct {
	channel *amqp.Channel

	Name         string
	exchangeType string

	durable    bool
	autoDelete bool
	internal   bool
	noWait     bool
}

func NewExchange(channel *amqp.Channel, name, exchangeType string) *Exchange {
	return &Exchange{
		channel: channel,

		Name:         name,
		exchangeType: exchangeType,
	}
}

func (exchange *Exchange) Durable() *Exchange {
	exchange.durable = true
	return exchange
}

func (exchange *Exchange) AutoDelete() *Exchange {
	exchange.autoDelete = true
	return exchange
}

func (exchange *Exchange) Internal() *Exchange {
	exchange.internal = true
	return exchange
}

func (exchange *Exchange) NoWait() *Exchange {
	exchange.noWait = true
	return exchange
}

func (exchange *Exchange) Declare() (*Exchange, error) {
	if err := exchange.channel.ExchangeDeclare(
		exchange.Name,
		exchange.exchangeType,
		exchange.durable,
		exchange.autoDelete,
		exchange.internal,
		exchange.noWait,
		nil,
	); err != nil {
		return exchange, err
	}

	return exchange, nil
}
