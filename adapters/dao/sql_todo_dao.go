package dao

import (
	"go-cleanarchitecture/domains"
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/domains/models/todo"

	"gorm.io/gorm"
)

type TodoDao SQLDao

var _ domains.TodoRepository = TodoDao{}

func NewSQLTodoDao(tt txType) (TodoDao, error) {
	dao, err := newSQLDao("todo", tt)
	return TodoDao(dao), err
}

func (dao TodoDao) Close() {
	dao.Close()
}

type TodoDto struct {
	ID          string `gorm:"id"`
	Name        string `gorm:"name"`
	Description string `gorm:"description"`
	UserID      string `gorm:"user_id"`
}

func (TodoDto) TableName() string {
	return "todo"
}

func (dao TodoDao) Get(id todo.ID) (models.Todo, errors.Domain, bool) {
	var dto TodoDto

	query := dao.
		conn.
		First(&dto, "id = ?", id.String())

	empty := models.Todo{}

	if query.Error == gorm.ErrRecordNotFound {
		return empty, errors.None, false
	} else if query.Error != nil {
		return empty, errors.External(query.Error), false
	}

	_id, _ := todo.NewID(dto.ID)
	name, _ := todo.NewName(dto.Name)
	description, _ := todo.NewDescription(dto.Description)
	return models.BuildTodo(_id, name, description), errors.None, true
}

func (dao TodoDao) GetByName(name todo.Name) (models.Todo, errors.Domain, bool) {
	var dto TodoDto

	query := dao.
		conn.
		Where("name = ?", name.Value()).
		Find(&dto)

	empty := models.Todo{}

	if query.Error == gorm.ErrRecordNotFound {
		return empty, errors.None, false
	} else if query.Error != nil {
		return empty, errors.External(query.Error), false
	}

	id, _ := todo.NewID(dto.ID)
	_name, _ := todo.NewName(dto.Name)
	description, _ := todo.NewDescription(dto.Description)
	return models.BuildTodo(id, _name, description), errors.None, true
}

func (dao TodoDao) Store(todo models.Todo) errors.Domain {
	dto := TodoDto{
		ID:          todo.Id().String(),
		Name:        todo.Name().Value(),
		Description: todo.Description().Value(),
	}

	return errors.External(dao.conn.Create(&dto).Error)
}
