package rmq_autoconnector

import (
	"time"
)

func (service *Service) Run() {
	ticker := time.NewTicker(service.interval)

	for {
		<-ticker.C

		if service.event.IsClosed() {
			service.reconnectMessage()
			if err := service.event.Reconnect(); err != nil {
				service.errorHandler(err)
			}
		}
	}
}

func (service *Service) RunAsync() {
	go func() {
		service.Run()
	}()
}

func (service *Service) Interval(interval time.Duration) *Service {
	service.interval = interval
	return service
}

func (service *Service) ErrorHandler(handler func(err error)) *Service {
	service.errorHandler = handler
	return service
}

func (service *Service) ReconnectMessage(handler func()) *Service {
	service.reconnectMessage = handler
	return service
}
