package dao

import (
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/domains/models/todo"
)

type TodosDao SQLDao

func NewSQLTodosDao() (TodosDao, error) {
	err, dao := newSQLDao("todos")
	return TodosDao(dao), err
}

func (dao TodosDao) Close() {
	dao.Close()
}

func (dao TodosDao) Get() (models.Todos, errors.Domain, bool) {
	var dtos []todoDto

	query := dao.
		conn.
		Find(&dtos)

	empty := models.Todos{}

	if query.RecordNotFound() {
		// リスト取得系操作なので空配列の戻り値が正しいためtrueを返している
		return empty, errors.None, true
	} else if query.Error != nil {
		return empty, errors.External(query.Error), false
	}

	todos := []models.Todo{}
	for _, dto := range dtos {
		// すでに永続化されているtodo自体は作成時のバリデーションを経由しているため
		// ここではバリデーションエラーはでないことを期待するためエラーは無視している。
		id, _ := todo.NewId(dto.Id)
		name, _ := todo.NewName(dto.Name)
		description, _ := todo.NewDescription(dto.Description)
		todos = append(todos, models.BuildTodo(id, name, description))
	}

	return models.NewTodos(todos), errors.None, true
}
