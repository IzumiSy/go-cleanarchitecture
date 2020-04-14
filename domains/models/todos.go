package models

type Todos struct {
	// [TODOの集合を表現するファーストクラス・コレクション]

	value []Todo
}

func NewTodos(value []Todo) Todos {
	return Todos{value}
}

func (todos Todos) Value() []Todo {
	return todos.value
}

func EmptyTodos() Todos {
	return Todos{[]Todo{}}
}
