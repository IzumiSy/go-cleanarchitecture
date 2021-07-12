package dao

import (
	"go-cleanarchitecture/domains"
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/domains/models/todo"
)

type TodoDao SQLDao

var _ domains.TodoRepository = TodoDao{}

func NewSQLTodoDao(name string, tt txType) (TodoDao, error) {
	err, dao := newSQLDao("todo", tt)
	return TodoDao(dao), err
}

func (dao TodoDao) Close() {
	dao.Close()
}

type todoDto struct {
	Id          string `gorm:"id"`
	Name        string `gorm:"name"`
	Description string `gorm:"description"`
}

func (dao TodoDao) Get(id todo.Id) (models.Todo, errors.Domain, bool) {
	var todo models.Todo

	query := dao.
		conn.
		Find(&todo, id.String())

	empty := models.Todo{}

	// Errorよりも先にRecordNotFoundをチェックしないと
	// レコードが存在しないというErrorとしてハンドリングされてしまう
	if query.RecordNotFound() {
		return empty, errors.None, false
	} else if query.Error != nil {
		return empty, errors.External(query.Error), true
	}

	return todo, errors.None, true
}

func (dao TodoDao) GetByName(name todo.Name) (models.Todo, errors.Domain, bool) {
	var dto todoDto

	query := dao.
		conn.
		Where("name = ?", name.Value()).
		Find(&dto)

	empty := models.Todo{}

	// Errorよりも先にRecordNotFoundをチェックしないと
	// レコードが存在しないというErrorとしてハンドリングされてしまう
	if query.RecordNotFound() {
		return empty, errors.None, false
	} else if query.Error != nil {
		return empty, errors.External(query.Error), false
	}

	// すでに永続化されているtodo自体は作成時のバリデーションを経由しているため
	// ここではバリデーションエラーはでないことを期待するためエラーは無視している。
	id, _ := todo.NewId(dto.Id)
	_name, _ := todo.NewName(dto.Name)
	description, _ := todo.NewDescription(dto.Description)

	return models.BuildTodo(id, _name, description), errors.None, true
}

func (dao TodoDao) Store(todo models.Todo) errors.Domain {
	dto := todoDto{
		Id:          todo.Id().String(),
		Name:        todo.Name().Value(),
		Description: todo.Description().Value(),
	}

	return errors.External(dao.conn.Create(&dto).Error)
}
