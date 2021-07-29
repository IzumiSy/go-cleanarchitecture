package dao

import (
	"go-cleanarchitecture/domains"
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/domains/models/todo"
	"go-cleanarchitecture/domains/models/user"
)

type MockTodoDao struct {
	GetResult       func() (models.Todo, errors.Domain, bool)
	GetByNameResult func() (models.Todo, errors.Domain, bool)
	StoreResult     func() errors.Domain
}

var _ domains.TodoRepository = MockTodoDao{}

func (m MockTodoDao) Get(id todo.ID) (models.Todo, errors.Domain, bool) {
	t, err, e := m.GetResult()
	return t, err, e
}

func (m MockTodoDao) GetByName(userID user.ID, name todo.Name) (models.Todo, errors.Domain, bool) {
	t, err, e := m.GetByNameResult()
	return t, err, e
}

func (m MockTodoDao) Store(todo models.Todo) errors.Domain {
	return m.StoreResult()
}
