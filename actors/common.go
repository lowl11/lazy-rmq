package actors

import (
	"context"
	"time"
)

func _ctx() (context.Context, func()) {
	return context.WithTimeout(context.Background(), time.Second*5)
}
