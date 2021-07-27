package domains

import "fmt"

type DomainEventID struct {
	Name string
}

type Entity interface {
	ID() fmt.Stringer
}

var (
	UserSignedUp      DomainEventID = DomainEventID{"user_singed_up"}
	UserAuthenticated DomainEventID = DomainEventID{"user_authenticated"}
	TodoCreated       DomainEventID = DomainEventID{"todo_created"}
)

type DomainEvent struct {
	ID     DomainEventID
	Entity Entity
}

type Pubsub interface {
	Publish(event DomainEvent) error
}
