package actors

import (
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	defaultContentType = "text/plain"
)

type Publisher struct {
	channel *amqp.Channel

	exchange  string
	routeKey  string
	mandatory bool

	contentType     string
	contentEncoding string
	deliveryMode    int
	priority        int
	replyTo         string
	expiration      string
}

func NewPublisher(channel *amqp.Channel) *Publisher {
	return &Publisher{
		channel:     channel,
		contentType: defaultContentType,
	}
}

func (publisher *Publisher) Exchange(exchange string) *Publisher {
	publisher.exchange = exchange
	return publisher
}

func (publisher *Publisher) Route(routeKey string) *Publisher {
	publisher.routeKey = routeKey
	return publisher
}

func (publisher *Publisher) Mandatory() *Publisher {
	publisher.mandatory = true
	return publisher
}

func (publisher *Publisher) ContentType(contentType string) *Publisher {
	if contentType == "" {
		return publisher
	}

	publisher.contentType = contentType
	return publisher
}

func (publisher *Publisher) DeliveryMode(mode int) *Publisher {
	publisher.deliveryMode = mode
	return publisher
}

func (publisher *Publisher) Priority(priority int) *Publisher {
	publisher.priority = priority
	return publisher
}

func (publisher *Publisher) ReplyTo(replyTo string) *Publisher {
	publisher.replyTo = replyTo
	return publisher
}

func (publisher *Publisher) Expiration(timeInSeconds string) *Publisher {
	publisher.expiration = timeInSeconds
	return publisher
}

func (publisher *Publisher) Encoding(mode string) *Publisher {
	publisher.contentEncoding = mode
	return publisher
}

func (publisher *Publisher) Send(body any) error {
	ctx, cancel := _ctx()
	defer cancel()

	bodyInBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	if err = publisher.channel.PublishWithContext(
		ctx,
		publisher.exchange,
		publisher.routeKey,
		publisher.mandatory,
		false, // always false, because parameter is deprecated
		amqp.Publishing{
			Body:            bodyInBytes,
			ContentType:     publisher.contentType,
			ContentEncoding: publisher.contentEncoding,
			DeliveryMode:    uint8(publisher.deliveryMode),
			Priority:        uint8(publisher.priority),
			Expiration:      publisher.expiration,
		},
	); err != nil {
		return err
	}

	return nil
}
