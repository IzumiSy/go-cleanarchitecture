package models

type Todos struct {
	value []Todo
}

func NewTodos(value []Todo) Todos {
	return Todos{value}
}

func EmptyTodos() Todos {
	return Todos{[]Todo{}}
}
