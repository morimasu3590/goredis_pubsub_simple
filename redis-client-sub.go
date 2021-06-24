package redis_pubsub_simple_example

import (
	"context"
	"fmt"

	_redis "github.com/go-redis/redis/v8"
)

type Redis struct {
	Connector     *_redis.Client
	PubSub        *_redis.PubSub
	Channel       string
	StringPattern string
}

var redis *Redis = nil
var ctx = context.Background()

func NewRedis(address string, password string) Redis {
	if redis == nil {
		redis = new(Redis)
		redis.Connector = _redis.NewClient(&_redis.Options{
			Addr:     address,
			Password: password,
			DB:       0,
		})
	}

	return *redis
}

func (redisPtr *Redis) RedisSUBSCRIBE() error {
	pubsub := redisPtr.Connector.Subscribe(ctx, redisPtr.Channel)
	redisPtr.PubSub = pubsub
	for {
		msgi, err := pubsub.Receive(ctx)
		if err != nil {
			return err
		}

		switch msg := msgi.(type) {
		case *_redis.Message:
			fmt.Println("Received", msg.Payload, "on Channel", msg.Channel)
		default:
			fmt.Println("Got control message", msg)
		}
	}

}

func (redisPtr *Redis) RedisPSUBSCRIBE() error {
	pubsub := redisPtr.Connector.PSubscribe(ctx, redisPtr.Channel, redisPtr.StringPattern)
	redisPtr.PubSub = pubsub
	for {
		msgi, err := pubsub.Receive(ctx)
		if err != nil {
			return err
		}

		switch msg := msgi.(type) {
		case *_redis.Message:
			fmt.Println("Received", msg.Payload, "on Channel", msg.Channel)
		default:
			fmt.Println("Got control message", msg)
		}
	}
}
