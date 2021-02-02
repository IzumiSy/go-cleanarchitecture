package dao

import (
	"go-cleanarchitecture/domains"
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/domains/models/todo"
)

type MockTodosDao struct{}

var _ domains.TodosRepository = MockTodosDao{}

func (_ MockTodosDao) GetByIds(ids []todo.Id) (models.Todos, errors.Domain) {
    return models.Todos{}, errors.Domain{}
}

func (_ MockTodosDao) Get() (models.Todos, errors.Domain) {
    return models.Todos{}, errors.Domain{}
}
