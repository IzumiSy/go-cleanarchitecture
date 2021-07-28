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

	return errors.Postconditional(adapter.conn.Send(string(event.Name()), eventBytes))
}

func (adapter RedisAdapter) RegisterSubscriber(eventID domains.EventName, subscriber func(payload []byte) error) {
	adapter.subscribers[string(eventID)] = subscriber
	if err := adapter.psc.Subscribe(string(eventID)); err != nil {
		adapter.logger.Error(fmt.Sprintf("Failed subscribing %s: %s", string(eventID), err.Error()))
	}
}

func (adapter RedisAdapter) Listen() {
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
