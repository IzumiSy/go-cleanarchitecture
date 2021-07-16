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

func (_ MockTodosDao) GetByIDs(ids []todo.ID) (models.Todos, errors.Domain) {
	return models.Todos{}, errors.Domain{}
}

func (_ MockTodosDao) GetByUserID(userId user.ID) (models.Todos, errors.Domain) {
	return models.Todos{}, errors.Domain{}
}
