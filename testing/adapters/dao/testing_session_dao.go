package dao

import (
	"go-cleanarchitecture/domains"
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/domains/models/session"
)

type mockSessionDao struct {
	StoreResult func() errors.Domain
	GetResult   func() (models.Session, errors.Domain, bool)
}

func NewMockSessionDao() mockSessionDao {
	return mockSessionDao{
		StoreResult: func() errors.Domain {
			return errors.None
		},
		GetResult: func() (models.Session, errors.Domain, bool) {
			return models.Session{}, errors.None, false
		},
	}
}

var _ domains.SessionRepository = mockSessionDao{}

func (m mockSessionDao) Store(session models.Session) errors.Domain {
	return m.StoreResult()
}

func (m mockSessionDao) Get(id session.ID) (models.Session, errors.Domain, bool) {
	s, err, e := m.GetResult()
	return s, err, e
}
