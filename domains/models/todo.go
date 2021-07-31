package models

import (
	"go-cleanarchitecture/domains/models/todo"
	"go-cleanarchitecture/domains/models/user"
)

type Todo struct {
	// [TODOを表現するエンティティ]

	id          todo.ID
	name        todo.Name
	description todo.Description

	userID      user.ID
	categoryIDs todo.CategoryIDs
}

// アプリケーション上における新規のTODOを作成する関数
// IDは内部で生成するためこの関数では外部から入力を受け付けない
func NewTodo(name todo.Name, description todo.Description, userID user.ID) Todo {
	return Todo{
		id:          todo.GenerateID(),
		name:        name,
		description: description,
		userID:      userID,
		categoryIDs: todo.EmptyCategoryIds(),
	}
}

// repositoryやfactory経由の生成において使用する関数
// 生成時のバリデーションをしないことに注意
func BuildTodo(id todo.ID, name todo.Name, description todo.Description, userID user.ID) Todo {
	return Todo{
		id:          id,
		name:        name,
		userID:      userID,
		description: description,
	}
}

func (todo Todo) ID() todo.ID {
	return todo.id
}

func (todo Todo) UserID() user.ID {
	return todo.userID
}

func (todo Todo) Name() todo.Name {
	return todo.name
}

func (todo Todo) Description() todo.Description {
	return todo.description
}

func (todo Todo) UpdateCategoryIds(categoryIds todo.CategoryIDs) {
	todo.categoryIDs = categoryIds
}
