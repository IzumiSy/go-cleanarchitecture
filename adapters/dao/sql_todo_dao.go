package dao

import (
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

func (dao TodoDao) Get(id todo.Id) (error, models.Todo) {
	var todo models.Todo

	err := dao.
		conn.
		Find(&todo, id.String()).
		Error

	return err, todo
}

func (dao TodoDao) GetByName(name todo.Name) (error, models.Todo) {
	var todo models.Todo

	err := dao.
		conn.
		Where("name = ?", name.Value()).
		Find(&todo).
		Error

	return err, todo
}

func (dao TodoDao) Store(todo models.Todo) error {
	return dao.
		conn.
		Create(&todo).
		Error
}
