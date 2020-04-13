package usecases

import (
	"go-cleanarchitecture/domains"
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/domains/models/todo"
)

type CreateTodoOutputPort interface {
	domains.OutputPort
	Write(todo models.Todo)
}

type CreateTodoParam struct {
	Name        string
	Description string
}

type createTodoUsecase struct {
	outputPort CreateTodoOutputPort
	todoDao    domains.TodoRepository
	logger     domains.Logger
}

func NewCreateTodoUsecase(
	outputPort CreateTodoOutputPort,
	todoDao domains.TodoRepository,
	logger domains.Logger,
) createTodoUsecase {
	return createTodoUsecase{outputPort, todoDao, logger}
}

func (usecase createTodoUsecase) Execute(params CreateTodoParam) {
	// [TODO作成を行うユースケース]
	// バリデーションルールは以下
	// - すでに同名のTODOが存在している場合にはTODOは作成できない

	var (
		NAME_INVALID = errors.Invalid("Name must not be duplicated")
	)

	name, err := todo.NewName(params.Name)
	if err.NotNil() {
		usecase.outputPort.Raise(err)
		return
	}

	description, err := todo.NewDescription(params.Description)
	if err.NotNil() {
		usecase.outputPort.Raise(err)
		return
	}

	currentTodo, err, exists := usecase.todoDao.GetByName(name)
	if err.NotNil() {
		usecase.outputPort.Raise(err)
		return
	}

	if exists {
		if currentTodo.Name() == name {
			usecase.outputPort.Raise(NAME_INVALID)
			return
		}
	}

	newTodo := models.NewTodo(name, description)
	err = usecase.todoDao.Store(newTodo)
	if err.NotNil() {
		usecase.outputPort.Raise(err)
		return
	}

	usecase.outputPort.Write(newTodo)
}
