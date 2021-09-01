package dao

import (
	"go-cleanarchitecture/domains"
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/domains/models/todo"
	"go-cleanarchitecture/domains/models/user"
)

type mockTodoDao struct {
	GetResult       func() (models.Todo, errors.Domain, bool)
	GetByNameResult func() (models.Todo, errors.Domain, bool)
	StoreResult     func() errors.Domain
}

var _ domains.TodoRepository = mockTodoDao{}

func NewMockTodoDao() mockTodoDao {
	return mockTodoDao{
		GetResult: func() (models.Todo, errors.Domain, bool) {
			return models.Todo{}, errors.None, false
		},
		GetByNameResult: func() (models.Todo, errors.Domain, bool) {
			return models.Todo{}, errors.None, false
		},
		StoreResult: func() errors.Domain {
			return errors.None
		},
	}
}

func (m mockTodoDao) Get(id todo.ID) (models.Todo, errors.Domain, bool) {
	t, err, e := m.GetResult()
	return t, err, e
}

func (m mockTodoDao) GetByName(userID user.ID, name todo.Name) (models.Todo, errors.Domain, bool) {
	t, err, e := m.GetByNameResult()
	return t, err, e
}

func (m mockTodoDao) Store(todo models.Todo) errors.Domain {
	return m.StoreResult()
}
