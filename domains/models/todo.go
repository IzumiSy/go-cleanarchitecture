package models

type Todo struct {
	id   string
	name string
}

// TODO needs a validation here
func NewTodo(name string) Todo {
	return Todo{name: name}
}

func (todo Todo) Name() string {
	return todo.name
}

func (todo Todo) Id() string {
	return todo.id
}
