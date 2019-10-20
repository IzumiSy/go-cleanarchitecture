package dao

import (
	"go-cleanarchitecture/domains/models"

	"github.com/jinzhu/gorm"
)

type TodoDao struct{}

func NewSqlTodoDao() TodoDao {
	return TodoDao{}
}

func (TodoDao) GetAll() ([]models.Todo, error) {
	var todos []models.Todo

	err := WithConnection(func(db *gorm.DB) error {
		return db.Find(&todos).Error
	})

	return todos, err
}

func (TodoDao) CreateOne(todo models.Todo) error {
	return WithConnection(func(db *gorm.DB) error {
		return db.Create(&todo).Error
	})
}
