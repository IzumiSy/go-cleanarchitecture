package models

import (
	"go-cleanarchitecture/domains/models/entity"
	"go-cleanarchitecture/domains/models/todo"
	"go-cleanarchitecture/domains/models/user"
)

type Todo struct {
	// [TODOを表現するエンティティ]
	// 集約ルートでありTodoHistoryエンティティを持つ

	id          todo.ID
	userID      user.ID
	name        todo.Name
	description todo.Description

	categoryIDs *todo.CategoryIDs
	histories   *todo.Histories
}

// アプリケーション上における新規のTODOを作成する関数
// IDは内部で生成するためこの関数では外部から入力を受け付けない
func NewTodo(name todo.Name, description todo.Description, userID user.ID) Todo {
	todoID := todo.ID{ID_: entity.GenerateID()}

	return Todo{
		id:          todoID,
		userID:      userID,
		name:        name,
		description: description,
		categoryIDs: todo.EmptyCategoryIds(),
		histories:   todo.NewHistories(todoID),
	}
}

// repositoryやfactory経由の生成において使用する関数
// 生成時のバリデーションをしないことに注意
func BuildTodo(
	id todo.ID,
	userID user.ID,
	name todo.Name,
	description todo.Description,
) Todo {
	return Todo{
		id:          id,
		userID:      userID,
		name:        name,
		description: description,
	}
}

func (t *Todo) ID() todo.ID {
	return t.id
}

func (t *Todo) UserID() user.ID {
	return t.userID
}

func (t *Todo) Name() todo.Name {
	return t.name
}

func (t *Todo) Description() todo.Description {
	return t.description
}

func (t *Todo) Histories() *todo.Histories {
	return t.histories
}

func (t *Todo) UpdateName(name todo.Name) *Todo {
	t.name = name
	t.histories.AddHistory(todo.NameUpdated, name)
	return t
}

func (t *Todo) UpdateDescription(description todo.Description) *Todo {
	t.description = description
	t.histories.AddHistory(todo.DescriptionUpdated, description)
	return t
}
