package usecases

import (
	"go-cleanarchitecture/domains"
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/domains/models/todo"
	"go-cleanarchitecture/domains/models/user"
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
}

func (uc CreateTodoUsecase) Build(params CreateTodoParam) domains.AuthorizedUsecase {
	return domains.NewAuthorizedUsecase(uc.OutputPort, func(currentUserID user.ID) {
		// [TODO作成を行うユースケース]
		// バリデーションルールは以下
		// - すでに同名のTODOが存在している場合にはTODOは作成できない
		// - 新しく作成できるTODOは100件まで

		var (
			NAME_INVALID          = errors.Invalid("Name must not be duplicated")
			MAXIMUM_TODOS_REACHED = errors.Invalid("Maximum TODOs reached")
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

		userID, err := user.NewID("d70f4845-b645-4271-bea9-3d5705e79e87")
		if err.NotNil() {
			uc.Logger.Warn(err.Error())
			uc.OutputPort.Raise(err)
			return
		}

		todos, err := uc.TodosDao.GetByUserID(userID)
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

		uc.OutputPort.Write(newTodo)
	})
}
