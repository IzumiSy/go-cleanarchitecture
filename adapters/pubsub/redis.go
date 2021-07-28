package pubsub

import (
	"encoding/json"
	"fmt"
	"go-cleanarchitecture/domains"
	"go-cleanarchitecture/domains/errors"

	"github.com/gomodule/redigo/redis"
)

type Subscriber = func(payload []byte) error

type RedisAdapter struct {
	conn        redis.Conn
	psc         redis.PubSubConn
	subscribers map[string]Subscriber
	logger      domains.Logger
}

func NewRedisAdapter(logger domains.Logger) (error, RedisAdapter) {
	conn, err := redis.Dial("tcp", "redis:6379")
	if err != nil {
		return err, RedisAdapter{}
	}

	return nil, RedisAdapter{
		conn:        conn,
		psc:         redis.PubSubConn{Conn: conn},
		subscribers: map[string]Subscriber{},
		logger:      logger,
	}
}

func (adapter RedisAdapter) Publish(event domains.Event) errors.Domain {
	eventBytes, err := json.Marshal(event)
	if err != nil {
		return errors.Postconditional(err)
	}

	_, err = adapter.conn.Do("PUBLISH", string(event.Name()), eventBytes)
	return errors.Postconditional(err)
}

func (adapter RedisAdapter) RegisterSubscriber(eventID domains.EventName, subscriber func(payload []byte) error) {
	adapter.subscribers[string(eventID)] = subscriber
}

func (adapter RedisAdapter) Listen() {
	var channels []string
	for c := range adapter.subscribers {
		channels = append(channels, c)
	}

	if err := adapter.psc.Subscribe(redis.Args{}.AddFlat(channels)); err != nil {
		adapter.logger.Error(fmt.Sprintf("Failed to start listening: %s", err.Error()))
	}

	for {
		switch n := adapter.psc.Receive().(type) {
		case error:
			adapter.logger.Error(fmt.Sprintf("Error listening subscribers: %s", n.Error()))
			return
		case redis.Message:
			subscriber, ok := adapter.subscribers[n.Channel]
			if ok {
				subscriber(n.Data)
			}
		case redis.Subscription:
			switch n.Kind {
			case "subscribe":
				adapter.logger.Info(fmt.Sprintf("%s subscribed", n.Channel))
			case "unsubscribed":
				adapter.logger.Info(fmt.Sprintf("%s unsubscribed", n.Channel))
			}
		}
	}
}
