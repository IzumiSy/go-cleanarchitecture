package dao

import (
	"go-cleanarchitecture/domains"
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/domains/models/authentication"
)

type mockAuthenticationDao struct {
	StoreResult      func() errors.Domain
	GetByEmailResult func() (models.Authentication, errors.Domain, bool)
}

func NewMockAuthenticationDao() mockAuthenticationDao {
	return mockAuthenticationDao{
		StoreResult: func() errors.Domain {
			return errors.None
		},
		GetByEmailResult: func() (models.Authentication, errors.Domain, bool) {
			return models.Authentication{}, errors.None, false
		},
	}
}

var _ domains.AuthenticationRepository = mockAuthenticationDao{}

func (m mockAuthenticationDao) Store(auth models.Authentication) errors.Domain {
	return m.StoreResult()
}

func (m mockAuthenticationDao) GetByEmail(email authentication.Email) (models.Authentication, errors.Domain, bool) {
	a, err, e := m.GetByEmailResult()
	return a, err, e
}
