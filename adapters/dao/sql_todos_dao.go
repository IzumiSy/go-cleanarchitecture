package dao

import (
	"go-cleanarchitecture/domains/models"
)

type TodosDao SQLDao

func NewSQLTodosDao() (TodosDao, error) {
	err, dao := newSQLDao()
	return TodosDao(dao), err
}

func (dao TodosDao) Close() {
	dao.Close()
}

func (dao TodosDao) Get() (error, models.Todos) {
	var todos []models.Todo

	err := dao.
		conn.
		Find(&todos).
		Error

	return err, models.NewTodos(todos)
}
