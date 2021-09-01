package adapters

import (
	"context"
	"encoding/json"
	"go-cleanarchitecture/domains"
	"go-cleanarchitecture/domains/usecases"
)

type Subscriber = func(payload []byte) error

func signedUpSubscriber(ctx context.Context, logger domains.Logger) Subscriber {
	return func(payload []byte) error {
		var event usecases.UserSignedUpEvent
		if err := json.Unmarshal(payload, &event); err != nil {
			return err
		}

		logger.Infof(ctx, "UserSignedUp event received: %s", event.CreatedAt)
		return nil
	}
}

func userAuthenticatedSubscriber(ctx context.Context, logger domains.Logger) Subscriber {
	return func(payload []byte) error {
		var event usecases.UserAuthenticatedEvent
		if err := json.Unmarshal(payload, &event); err != nil {
			return err
		}

		logger.Infof(ctx, "UserAuthenticated event received: %s", event.CreatedAt)
		return nil
	}
}

func todoCreatedSubscriber(ctx context.Context, logger domains.Logger) Subscriber {
	return func(payload []byte) error {
		var event usecases.TodoCreatedEvent
		if err := json.Unmarshal(payload, &event); err != nil {
			return err
		}

		logger.Infof(ctx, "TodoCreated event received: %s", event.CreatedAt)
		return nil
	}
}
