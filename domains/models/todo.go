package models

import (
	"go-cleanarchitecture/domains/models/todo"
)

type Todo struct {
	// [TODOを表現するエンティティ]

	id          todo.Id
	name        todo.Name
	description todo.Description
}

// アプリケーション上における新規のTODOを作成する関数
// IDは内部で生成するためこの関数では外部から入力を受け付けない
func NewTodo(name todo.Name, description todo.Description) Todo {
	return Todo{
		id:          todo.GenerateId(),
		name:        name,
		description: description,
	}
}

// repositoryやfactory経由の生成において使用する関数
// 生成時のバリデーションをしないことに注意
func BuildTodo(id todo.Id, name todo.Name, description todo.Description) Todo {
	return Todo{
		id:          id,
		name:        name,
		description: description,
	}
}

func (todo Todo) Id() todo.Id {
	return todo.id
}

func (todo Todo) Name() todo.Name {
	return todo.name
}

func (todo Todo) Description() todo.Description {
	return todo.description
}
