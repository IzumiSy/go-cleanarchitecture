package dao

import (
	"context"
	"go-cleanarchitecture/domains"
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/domains/models/entity"
	"go-cleanarchitecture/domains/models/todo"
	"go-cleanarchitecture/domains/models/user"

	"gorm.io/gorm"
)

type TodosDao SQLDao

var _ domains.TodosRepository = TodosDao{}

func (driver Driver) NewSQLTodosDao(tt txType) (TodosDao, error) {
	dao, err := driver.newSQLDao("todo", tt)
	return TodosDao(dao), err
}

func (dao TodosDao) Close() {
	dao.Close()
}

func (dao TodosDao) GetByIDs(ids []todo.ID) (models.Todos, errors.Domain) {
	return models.EmptyTodos(), errors.None // todo: あとで実装する
}

func (dao TodosDao) GetByUserID(userID user.ID) (models.Todos, errors.Domain) {
	var dtos []TodoDto

	query := dao.
		conn.
		WithContext(context.Background()).
		Where("user_id = ?", userID.String()).
		Find(&dtos)

	empty := models.Todos{}

	if query.Error == gorm.ErrRecordNotFound {
		return empty, errors.None
	} else if query.Error != nil {
		return empty, errors.Postconditional(query.Error)
	}

	todos := []models.Todo{}
	for _, dto := range dtos {
		// 永続化済みのデータの取り出しでバリデーションエラーはないはずなのでエラーは無視する
		id_, _ := entity.NewID(dto.ID)
		name, _ := todo.NewName(dto.Name)
		description, _ := todo.NewDescription(dto.Description)
		id := todo.ID{ID_: id_}
		todos = append(todos, models.BuildTodo(id, userID, name, description))
	}

	return models.NewTodos(todos), errors.None
}
