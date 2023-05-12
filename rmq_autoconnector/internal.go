package rmq_autoconnector

import "time"

const (
	defaultInterval = time.Second * 60
)

var (
	defaultErrorHandler     = func(err error) {}
	defaultReconnectMessage = func() {}
)
