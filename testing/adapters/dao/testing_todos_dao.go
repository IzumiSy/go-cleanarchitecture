package dao

import (
	"go-cleanarchitecture/domains"
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/domains/models/todo"
	"go-cleanarchitecture/domains/models/user"
)

type MockTodosDao struct{}

var _ domains.TodosRepository = MockTodosDao{}

func (_ MockTodosDao) GetByIDs(ids []todo.Id) (models.Todos, errors.Domain) {
	return models.Todos{}, errors.Domain{}
}

func (_ MockTodosDao) GetByUserID(userId user.Id) (models.Todos, errors.Domain) {
	return models.Todos{}, errors.Domain{}
}
