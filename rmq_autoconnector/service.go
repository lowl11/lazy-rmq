package rmq_autoconnector

import (
	"github.com/lowl11/lazy-rmq/rabbit_event"
	"time"
)

type Service struct {
	event            *rabbit_event.Event
	interval         time.Duration
	errorHandler     func(err error)
	reconnectMessage func()
}

func New(event *rabbit_event.Event) *Service {
	return &Service{
		event:            event,
		interval:         defaultInterval,
		errorHandler:     defaultErrorHandler,
		reconnectMessage: defaultReconnectMessage,
	}
}
