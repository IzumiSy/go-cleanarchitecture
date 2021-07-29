package adapters

import (
	"go-cleanarchitecture/domains"
	"go-cleanarchitecture/domains/errors"
)

type MockPublisher struct{}

func (MockPublisher) Publish(event domains.Event) errors.Domain {
	return errors.None
}
