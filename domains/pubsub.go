package domains

import (
	"go-cleanarchitecture/domains/errors"

	"github.com/google/uuid"
)

type EventName string

var (
	UserSignedUp      EventName = EventName("user_singed_up")
	UserAuthenticated EventName = EventName("user_authenticated")
	TodoCreated       EventName = EventName("todo_created")
)

type EventID uuid.UUID

func NewEventID() EventID {
	return EventID(uuid.New())
}

type Event interface {
	Name() EventName
	ID() EventID
}

type EventPublisher interface {
	Publish(event Event) errors.Domain
}
