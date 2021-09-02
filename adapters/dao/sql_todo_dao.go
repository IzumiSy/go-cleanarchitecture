package dao

import (
	"context"
	"go-cleanarchitecture/domains"
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/domains/models/entity"
	"go-cleanarchitecture/domains/models/todo"
	"go-cleanarchitecture/domains/models/user"

	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TodoDao SQLDao

var _ domains.TodoRepository = TodoDao{}

func (driver Driver) NewSQLTodoDao(tt txType) (TodoDao, error) {
	dao, err := driver.newSQLDao("todo", tt)
	return TodoDao(dao), err
}

func (dao TodoDao) Close() {
	dao.Close()
}

type TodoDto struct {
	ID          string `gorm:"column:id"`
	UserID      string `gorm:"column:user_id;not null;index;index:idx_user_id_name"`
	Name        string `gorm:"column:name;not null;index:idx_user_id_name"`
	Description string `gorm:"column:description;not null"`

	Categories []TodoCategoryDto `gorm:"many2many:todo_categories;"`
	Histories  []TodoHistoryDto  `gorm:"foreignKey:todo_id;constraint:OnDelete:CASCADE"`
}

func (TodoDto) TableName() string {
	return "todo"
}

type TodoCategoryDto struct {
	Name   string `gorm:"column:name;primaryKey"`
	UserID string `gorm:"column:user_id;primaryKey"`

	Todos []TodoDto `gorm:"many2many:todo_categories;"`
}

func (TodoCategoryDto) TableName() string {
	return "todo_category"
}

type TodoHistoryDto struct {
	ID          string    `gorm:"column:id"`
	TodoID      string    `gorm:"column:todo_id;not null;index"`
	HistoryType string    `gorm:"column:history_type;not null"`
	Previous    *string   `gorm:"column:previous"`
	Current     *string   `gorm:"column:current"`
	CreatedAt   time.Time `gorm:"column:created_at;not null"`
}

func (TodoHistoryDto) TableName() string {
	return "todo_history"
}

func (dao TodoDao) Get(id todo.ID) (models.Todo, errors.Domain, bool) {
	var dto TodoDto

	query := dao.
		conn.
		WithContext(context.Background()).
		Preload("Histories").
		Preload("Categories").
		Take(&dto, "id = ?", id.String())

	empty := models.Todo{}

	if query.Error == gorm.ErrRecordNotFound {
		return empty, errors.None, false
	} else if query.Error != nil {
		return empty, errors.Postconditional(query.Error), false
	}

	todoID_, _ := entity.NewID(dto.ID)
	name, _ := todo.NewName(dto.Name)
	description, _ := todo.NewDescription(dto.Description)
	userID_, _ := entity.NewID(dto.UserID)
	todoID := todo.ID{ID_: todoID_}
	userID := user.ID{ID_: userID_}
	return models.BuildTodo(todoID, userID, name, description), errors.None, true
}

func (dao TodoDao) GetByName(userID user.ID, name todo.Name) (models.Todo, errors.Domain, bool) {
	var dto TodoDto

	query := dao.
		conn.
		WithContext(context.Background()).
		Preload("Histories").
		Preload("Categories").
		Where("user_id = ?", userID.String()).
		Where("name = ?", name.Value()).
		Take(&dto)

	empty := models.Todo{}

	if query.Error == gorm.ErrRecordNotFound {
		return empty, errors.None, false
	} else if query.Error != nil {
		return empty, errors.Postconditional(query.Error), false
	}

	todoID_, _ := entity.NewID(dto.ID)
	description, _ := todo.NewDescription(dto.Description)
	todoID := todo.ID{ID_: todoID_}
	return models.BuildTodo(todoID, userID, name, description), errors.None, true
}

func (dao TodoDao) Store(todo models.Todo) errors.Domain {
	var histDtos []TodoHistoryDto

	histories := todo.Histories()
	for i := 0; i < histories.Length(); i++ {
		if hist := histories.At(i); hist != nil {
			var previous *string = nil
			if hist.Previous() != nil {
				previous_ := hist.Previous().Value()
				previous = &previous_
			}

			var current *string = nil
			if hist.Current() != nil {
				current_ := hist.Current().Value()
				current = &current_
			}

			histDtos = append(histDtos, TodoHistoryDto{
				ID:          (*hist).ID().String(),
				HistoryType: (*hist).Type().String(),
				Previous:    previous,
				Current:     current,
				CreatedAt:   hist.CreatedAt().Value(),
			})
		}
	}

	dto := TodoDto{
		ID:          todo.ID().String(),
		UserID:      todo.UserID().String(),
		Name:        todo.Name().Value(),
		Description: todo.Description().Value(),
		Categories:  []TodoCategoryDto{},
		Histories:   histDtos,
	}

	return errors.Postconditional(
		dao.
			conn.
			WithContext(context.Background()).
			Clauses(clause.OnConflict{UpdateAll: true}).
			Create(&dto).
			Error,
	)
}
