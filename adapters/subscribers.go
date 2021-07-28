package adapters

import (
	"encoding/json"
	"fmt"
	"go-cleanarchitecture/adapters/loggers"
	"go-cleanarchitecture/domains/usecases"
)

func signedUpHandler(payload []byte) error {
	var event usecases.UserSignedUpEvent
	if err := json.Unmarshal(payload, &event); err != nil {
		return err
	}

	logger, err := loggers.NewZapLogger("config/zap.json")
	if err != nil {
		return err
	}

	logger.Info(fmt.Sprintf("UserSignedUp event received: %s", event.CreatedAt))
	return nil
}

func userAuthenticatedHandler(payload []byte) error {
	var event usecases.UserAuthenticatedEvent
	if err := json.Unmarshal(payload, &event); err != nil {
		return err
	}

	logger, err := loggers.NewZapLogger("config/zap.json")
	if err != nil {
		return err
	}

	logger.Info(fmt.Sprintf("UserAuthenticated event received: %s", event.CreatedAt))
	return nil
}

func todoCreatedHandler(payload []byte) error {
	var event usecases.TodoCreatedEvent
	if err := json.Unmarshal(payload, &event); err != nil {
		return err
	}

	logger, err := loggers.NewZapLogger("config/zap.json")
	if err != nil {
		return err
	}

	logger.Info(fmt.Sprintf("TodoCreated event received: %s", event.CreatedAt))
	return nil
}
