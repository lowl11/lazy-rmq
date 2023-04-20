# lazy-rm

> simple wrapper library on AMQP 

### Usage example
<hr>

```go
package main

import (
	"fmt"
	"github.com/lowl11/lazy-rmq/exchanges"
	"github.com/lowl11/lazy-rmq/rabbit_event"
	"github.com/lowl11/lazy-rmq/rabbit_service"
	"log"
)

func main() {
	// create connections
	connection, err := rabbit_service.NewConnection("amqp://lazyuser:qwerty@127.0.0.1:5672")
	if err != nil {
		log.Fatal("rmq connection: ", err)
	}

	// creat event
	rabbit, err := rabbit_event.New(connection)
	if err != nil {
		log.Fatal("create event: ", err)
	}

	// declare queue
	var numbersQueue *rabbit_service.Queue
	var wordsQueue *rabbit_service.Queue

	if numbersQueue, err = rabbit.
		Queue("numbers").
		Declare(); err != nil {
		log.Fatal("declare queue: ", err)
	}

	if wordsQueue, err = rabbit.
		Queue("words").
		Declare(); err != nil {
		log.Fatal("declare queue: ", err)
	}

	// declare exchange
	var directExchange *rabbit_service.Exchange
	if directExchange, err = rabbit.
		Exchange("my_direct_exchange", exchanges.Direct).
		Declare(); err != nil {
		log.Fatal("declare exchange: ", err)
	}

	// bindings
	if err = rabbit.
		Bind(numbersQueue, directExchange).
		Declare(); err != nil {
		log.Fatal("declare binding: ", err)
	}

	if err = rabbit.
		Bind(wordsQueue, directExchange).
		Declare(); err != nil {
		log.Fatal("declare binding: ", err)
	}

	// consume messages
	consumer := rabbit.Consumer()

	if err = consumer.Handler(numbersQueue, func(ctx *rabbit_service.HandlerContext) error {
		//if string(ctx.Body) == "\"number #1\"" {
		//	return errors.New("erroooour")
		//}

		fmt.Printf("numbers: %s\n", ctx.Body)
		return nil
	}); err != nil {
		log.Fatal("start consuming err: ", err)
	}

	if err = consumer.Handler(wordsQueue, func(ctx *rabbit_service.HandlerContext) error {
		fmt.Printf("words: %s\n", ctx.Body)
		return nil
	}); err != nil {
		log.Fatal("start consuming err: ", err)
	}

	// publish message
	publisher := rabbit.Publisher().
		Exchange("my_direct_exchange").
		Route("numbers").
		ContentType("application/json").
		DeliveryMode(1)
	_ = publisher

	if err = publisher.Send("number #1"); err != nil {
		log.Fatal("send message error")
	}

	if err = publisher.Send("number #2"); err != nil {
		log.Fatal("send message error")
	}

	if err = publisher.Send("number #3"); err != nil {
		log.Fatal("send message error")
	}

	infinite := make(chan bool)
	<-infinite
}

```