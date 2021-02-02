package dao

import (
	"go-cleanarchitecture/domains"
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/domains/models/todo"
)

type MockTodoDao struct{}

var _ domains.TodoRepository = MockTodoDao{}

func (_ MockTodoDao) Get(id todo.Id) (models.Todo, errors.Domain, bool) {
	return models.Todo{}, errors.Domain{}, true
}

func (_ MockTodoDao) GetByName(name todo.Name) (models.Todo, errors.Domain, bool) {
	return models.Todo{}, errors.Domain{}, true
}

func (_ MockTodoDao) Store(todo models.Todo) errors.Domain {
    return errors.Domain{}
}
