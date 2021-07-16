package adapters

import "go-cleanarchitecture/domains/models"

type MockAuthorizer struct{}

func (a MockAuthorizer) Run() (models.Session, error) {
	return models.Session{}, nil
}
