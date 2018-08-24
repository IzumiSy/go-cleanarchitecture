package dao

// 本当はこれはinfrastructureパッケージなどに入れるほうがよいか

import (
	"github.com/IzumiSy/go-cleanarchitecture/domains/models"
)

type TodoDao interface {
	GetAll() ([]models.Todo, error)
	CreateOne(todo models.Todo) error
}
