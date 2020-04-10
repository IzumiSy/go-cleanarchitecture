package dao

import (
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/domains/models/todo"
)

type TodoDao SQLDao

func NewSQLTodoDao() (TodoDao, error) {
	err, dao := newSQLDao()
	return TodoDao(dao), err
}

func (dao TodoDao) Close() {
	dao.Close()
}

func (dao TodoDao) Get(id todo.Id) (models.Todo, errors.Domain) {
	var todo models.Todo

	err := dao.
		conn.
		Find(&todo, id.String()).
		Error

	return todo, errors.External(err)
}

func (dao TodoDao) GetByName(name todo.Name) (models.Todo, errors.Domain) {
	var todo models.Todo

	err := dao.
		conn.
		Where("name = ?", name.Value()).
		Find(&todo).
		Error

	return todo, errors.External(err)
}

func (dao TodoDao) Store(todo models.Todo) errors.Domain {
	return errors.External(
		dao.
			conn.
			Create(&todo).
			Error,
	)
}
