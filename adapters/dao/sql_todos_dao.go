package dao

import (
	"go-cleanarchitecture/domains"
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/domains/models/todo"
	"go-cleanarchitecture/domains/models/user"
)

type TodosDao SQLDao

var _ domains.TodosRepository = TodosDao{}

func NewSQLTodosDao(tt txType) (TodosDao, error) {
	err, dao := newSQLDao("todo", tt)
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

	if query.RecordNotFound() {
		// リスト取得系操作なので空配列の戻り値が正しいためtrueを返している
		return empty, errors.None
	} else if query.Error != nil {
		return empty, errors.External(query.Error)
	}

	todos := []models.Todo{}
	for _, dto := range dtos {
		// すでに永続化されているtodo自体は作成時のバリデーションを経由しているため
		// ここではバリデーションエラーはでないことを期待するためエラーは無視している。
		id, _ := todo.NewID(dto.ID)
		name, _ := todo.NewName(dto.Name)
		description, _ := todo.NewDescription(dto.Description)
		todos = append(todos, models.BuildTodo(id, name, description))
	}

	return models.NewTodos(todos), errors.None
}
