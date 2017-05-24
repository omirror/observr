package broker

import (
	"context"

	"github.com/thoas/observr/config"
)

type Event interface {
	Name() string
	ToBytes() ([]byte, error)
}

type Handler func(ctx context.Context, message []byte)

type Broker interface {
	Publish(ctx context.Context, event Event) error
	Run(ctx context.Context, handlers map[string]Handler) error
	Stop()
}

func Load(cfg config.Broker) (Broker, error) {
	return NewAMQPBroker(cfg)
}
