package repositories

import (
	"go-cleanarchitecture/domains/models"
)

type TodoRepository interface {
	GetAll() ([]models.Todo, error)
	CreateOne(todo models.Todo) error
}
