package dao

import (
	"go-cleanarchitecture/domains"
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/domains/models/todo"
	"go-cleanarchitecture/domains/models/user"
)

type mockTodosDao struct {
	GetByIDsResult    func() (models.Todos, errors.Domain)
	GetByUserIDResult func() (models.Todos, errors.Domain)
}

var _ domains.TodosRepository = mockTodosDao{}

func NewMockTodosDao() mockTodosDao {
	return mockTodosDao{
		GetByIDsResult: func() (models.Todos, errors.Domain) {
			return models.Todos{}, errors.None
		},
		GetByUserIDResult: func() (models.Todos, errors.Domain) {
			return models.Todos{}, errors.None
		},
	}
}

func (m mockTodosDao) GetByIDs(ids []todo.ID) (models.Todos, errors.Domain) {
	t, e := m.GetByIDsResult()
	return t, e
}

func (m mockTodosDao) GetByUserID(userId user.ID) (models.Todos, errors.Domain) {
	t, e := m.GetByUserIDResult()
	return t, e
}
