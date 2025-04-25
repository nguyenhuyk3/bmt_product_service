package services

import "context"

type IMessageBroker interface {
	SendMessage(ctx context.Context, topic, key string, message interface{}) error
	Close()
}
