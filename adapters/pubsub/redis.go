package pubsub

import (
	"go-cleanarchitecture/domains"

	"github.com/gomodule/redigo/redis"
)

type Subscriber = func(payload []byte) error

type RedisAdapter struct {
	conn        redis.Conn
	psc         redis.PubSubConn
	subscribers map[string]Subscriber
}

func NewRedisAdapter() (error, RedisAdapter) {
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		return err, RedisAdapter{}
	}

	return nil, RedisAdapter{
		conn:        conn,
		psc:         redis.PubSubConn{Conn: conn},
		subscribers: map[string]Subscriber{},
	}
}

func (adapter RedisAdapter) Publish(event domains.DomainEvent) error {
	return adapter.conn.Send(string(event.ID()), event)
}

func (adapter RedisAdapter) RegisterSubscriber(eventID domains.DomainEventID, subscriber func(payload []byte) error) {
	adapter.subscribers[string(eventID)] = subscriber
}

func (adapter RedisAdapter) Listen() {
	for {
		switch n := adapter.psc.Receive().(type) {
		case error:
			return
		case redis.Message:
			subscriber, ok := adapter.subscribers[n.Channel]
			if ok {
				subscriber(n.Data)
			}
		case redis.Subscription:
			// handling notification
		}
	}
}
