package dao

import (
	"go-cleanarchitecture/domains"
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/domains/models/todo"
	"go-cleanarchitecture/domains/models/user"

	"gorm.io/gorm"
)

type TodosDao SQLDao

var _ domains.TodosRepository = TodosDao{}

func NewSQLTodosDao(tt txType) (TodosDao, error) {
	dao, err := newSQLDao("todo", tt)
	return TodosDao(dao), err
}

func (dao TodosDao) Close() {
	dao.Close()
}

func (dao TodosDao) GetByIDs(ids []todo.ID) (models.Todos, errors.Domain) {
	return models.EmptyTodos(), errors.None // todo: あとで実装する
}

func (dao TodosDao) GetByUserID(userId user.ID) (models.Todos, errors.Domain) {
	var dtos []todoDto

	query := dao.
		conn.
		Where("user_id = ?", userId.String()).
		Find(&dtos)

	empty := models.Todos{}

	if query.Error == gorm.ErrRecordNotFound {
		return empty, errors.None
	} else if query.Error != nil {
		return empty, errors.External(query.Error)
	}

	todos := []models.Todo{}
	for _, dto := range dtos {
		// 永続化済みのデータの取り出しでバリデーションエラーはないはずなのでエラーは無視する
		id, _ := todo.NewID(dto.ID)
		name, _ := todo.NewName(dto.Name)
		description, _ := todo.NewDescription(dto.Description)
		todos = append(todos, models.BuildTodo(id, name, description))
	}

	return models.NewTodos(todos), errors.None
}
