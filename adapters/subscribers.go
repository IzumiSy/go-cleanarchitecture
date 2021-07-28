package adapters

import (
	"encoding/json"
	"fmt"
	"go-cleanarchitecture/domains"
	"go-cleanarchitecture/domains/usecases"
)

type Subscriber = func(payload []byte) error

func signedUpHandler(logger domains.Logger) Subscriber {
	return func(payload []byte) error {
		var event usecases.UserSignedUpEvent
		if err := json.Unmarshal(payload, &event); err != nil {
			return err
		}

		logger.Info(fmt.Sprintf("UserSignedUp event received: %s", event.CreatedAt))
		return nil
	}
}

func userAuthenticatedHandler(logger domains.Logger) Subscriber {
	return func(payload []byte) error {
		var event usecases.UserAuthenticatedEvent
		if err := json.Unmarshal(payload, &event); err != nil {
			return err
		}

		logger.Info(fmt.Sprintf("UserAuthenticated event received: %s", event.CreatedAt))
		return nil
	}
}

func todoCreatedHandler(logger domains.Logger) Subscriber {
	return func(payload []byte) error {
		var event usecases.TodoCreatedEvent
		if err := json.Unmarshal(payload, &event); err != nil {
			return err
		}

		logger.Info(fmt.Sprintf("TodoCreated event received: %s", event.CreatedAt))
		return nil
	}
}
