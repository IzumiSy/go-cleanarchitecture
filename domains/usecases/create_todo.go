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

func (uc CreateTodoUsecase) Build(params CreateTodoParam) domains.AuthorizedUsecase {
	return domains.NewAuthorizedUsecase(uc.OutputPort, func(currentUserID user.ID) {
		// [TODO作成を行うユースケース]
		// バリデーションルールは以下
		// - すでに同名のTODOが存在している場合にはTODOは作成できない
		// - 新しく作成できるTODOは100件まで

		var (
			NAME_INVALID          = errors.Preconditional("Name must not be duplicated")
			MAXIMUM_TODOS_REACHED = errors.Preconditional("Maximum TODOs reached")
		)

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

		currentTodo, err, exists := uc.TodoDao.GetByName(name)
		if err.NotNil() {
			uc.Logger.Error(err.Error())
			uc.OutputPort.Raise(err)
			return
		}

		if exists {
			if currentTodo.Name() == name {
				uc.OutputPort.Raise(NAME_INVALID)
				return
			}
		}

		todos, err := uc.TodosDao.GetByUserID(currentUserID)
		if todos.Size() > 100 {
			uc.OutputPort.Raise(MAXIMUM_TODOS_REACHED)
			return
		}

		newTodo := models.NewTodo(name, description)
		if err = uc.TodoDao.Store(newTodo); err.NotNil() {
			uc.Logger.Error(err.Error())
			uc.OutputPort.Raise(err)
			return
		}

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
