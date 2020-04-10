package usecases

import (
	"go-cleanarchitecture/domains"
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/domains/models/todo"
)

type CreateTodoOutputPort interface {
	OutputPort
	Write(todo models.Todo, result error)
}

type CreateTodoParam struct {
	Name        string
	Description string
}

type createTodoUsecase struct {
	outputPort CreateTodoOutputPort
	todoDao    domains.TodoRepository
}

func NewCreateTodoUsecase(outputPort CreateTodoOutputPort, todoDao domains.TodoRepository) createTodoUsecase {
	return createTodoUsecase{outputPort, todoDao}
}

func (usecase createTodoUsecase) Execute(params CreateTodoParam) {
	err, name := todo.NewName(params.Name)
	if err != nil {
		usecase.outputPort.Raise(err)
		return
	}

	err, description := todo.NewDescription(params.Description)
	if err != nil {
		usecase.outputPort.Raise(err)
		return
	}

	err, currentTodo := usecase.todoDao.GetByName(name)
	if err != nil {
		usecase.outputPort.Raise(err)
	}
	if currentTodo.Name() == name {
		// validation error
	}

	newTodo := models.NewTodo(name, description)
	err = usecase.todoDao.Store(newTodo)
	usecase.outputPort.Write(newTodo, err)
}
