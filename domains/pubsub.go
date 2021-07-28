package domains

type DomainEventID string

var (
	UserSignedUp      DomainEventID = DomainEventID("user_singed_up")
	UserAuthenticated DomainEventID = DomainEventID("user_authenticated")
	TodoCreated       DomainEventID = DomainEventID("todo_created")
)

type DomainEvent interface {
	Name() DomainEventID
}

type EventPublisher interface {
	Publish(event DomainEvent) error
}
