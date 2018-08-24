package models

type Todo struct {
	Id   string
	Name string
}

// TODO needs a validation here
func NewTodo(name string) Todo {
	return Todo{Name: name}
}
