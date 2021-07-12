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

type createTodoUsecase struct {
	outputPort CreateTodoOutputPort
	todoDao    domains.TodoRepository
	todosDao   domains.TodosRepository
	logger     domains.Logger
}

func NewCreateTodoUsecase(
	outputPort CreateTodoOutputPort,
	todoDao domains.TodoRepository,
	todosDao domains.TodosRepository,
	logger domains.Logger,
) createTodoUsecase {
	return createTodoUsecase{outputPort, todoDao, todosDao, logger}
}

func (usecase createTodoUsecase) Execute(params CreateTodoParam) {
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

	userID, err := user.New("user_id")
	if err.NotNil() {
		usecase.outputPort.Raise(err)
		return
	}

	todos, err := usecase.todosDao.GetByUserID(userID)
	if todos.Size() > 100 {
		usecase.outputPort.Raise(MAXIMUM_TODOS_REACHED)
		return
	}

	newTodo := models.NewTodo(name, description)
	err = usecase.todoDao.Store(newTodo)
	if err.NotNil() {
		usecase.outputPort.Raise(err)
		return
	}

	usecase.outputPort.Write(newTodo)
}
