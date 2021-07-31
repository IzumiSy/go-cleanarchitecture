package usecases

import (
	"fmt"
	"go-cleanarchitecture/domains"
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/domains/models/todo"
	"go-cleanarchitecture/domains/models/user"
	"time"
)

type CreateTodoOutputPort interface {
	domains.OutputPort
	Write(todo models.Todo)
}

type CreateTodoParam struct {
	Name        string
	Description string
}

type CreateTodoUsecase struct {
	OutputPort CreateTodoOutputPort
	TodoDao    domains.TodoRepository
	TodosDao   domains.TodosRepository
	Logger     domains.Logger
	Publisher  domains.EventPublisher
}

var (
	uc_TODO_NAME_NOT_UNIQUE  = errors.Preconditional("Name must be unique")
	uc_MAXIMUM_TODOS_REACHED = errors.Preconditional("Maximum TODOs reached")
)

func (uc CreateTodoUsecase) Build(params CreateTodoParam) domains.AuthorizedUsecase {
	return domains.NewAuthorizedUsecase(uc.OutputPort, func(currentUserID user.ID) {
		// [TODO作成を行うユースケース]
		// バリデーションルールは以下
		// - すでに同名のTODOが存在している場合にはTODOは作成できない
		// - 新しく作成できるTODOは100件まで

		name, err := todo.NewName(params.Name)
		if err.NotNil() {
			uc.Logger.Warn(err.Error())
			uc.OutputPort.Raise(err)
			return
		}

		description, err := todo.NewDescription(params.Description)
		if err.NotNil() {
			uc.Logger.Warn(err.Error())
			uc.OutputPort.Raise(err)
			return
		}

		currentTodo, err, exists := uc.TodoDao.GetByName(currentUserID, name)
		if err.NotNil() {
			uc.Logger.Error(err.Error())
			uc.OutputPort.Raise(err)
			return
		}

		if exists {
			if currentTodo.Name() == name {
				uc.Logger.Warn(fmt.Sprintf("Validation failed: %s", uc_TODO_NAME_NOT_UNIQUE.Error()))
				uc.OutputPort.Raise(uc_TODO_NAME_NOT_UNIQUE)
				return
			}
		}

		todos, err := uc.TodosDao.GetByUserID(currentUserID)
		if todos.Size() >= 100 {
			uc.Logger.Warn(fmt.Sprintf("Validation failed: %s", uc_MAXIMUM_TODOS_REACHED.Error()))
			uc.OutputPort.Raise(uc_MAXIMUM_TODOS_REACHED)
			return
		}

		newTodo := models.NewTodo(name, description)
		if err = uc.TodoDao.Store(newTodo); err.NotNil() {
			uc.Logger.Error(err.Error())
			uc.OutputPort.Raise(err)
			return
		}
		uc.Logger.Info(fmt.Sprintf("New todo stored: %s", newTodo.ID().String()))

		event := TodoCreatedEvent{
			TodoID:      newTodo.ID().String(),
			Name_:       newTodo.Name().Value(),
			Description: newTodo.Description().Value(),
			CreatedAt:   time.Now(),
		}
		if err := uc.Publisher.Publish(event); err.NotNil() {
			uc.Logger.Error(fmt.Sprintf("Failed publishing event: %s", err.Error()))
		}

		uc.Logger.Info(fmt.Sprintf("Event published: %s", event.ID()))
		uc.OutputPort.Write(newTodo)
	})
}

type TodoCreatedEvent struct {
	TodoID      string
	Name_       string
	Description string
	CreatedAt   time.Time
}

func (TodoCreatedEvent) Name() domains.EventName {
	return domains.TodoCreated
}

func (TodoCreatedEvent) ID() domains.EventID {
	return domains.NewEventID()
}
