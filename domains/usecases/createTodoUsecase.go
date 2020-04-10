package usecases

import (
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/domains/models/todo"
	"go-cleanarchitecture/domains/repositories"
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
	todoDao    repositories.TodoRepository
}

func NewCreateTodoUsecase(outputPort CreateTodoOutputPort, todoDao repositories.TodoRepository) createTodoUsecase {
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

	newTodo := models.NewTodo(name, description)
	err = usecase.todoDao.CreateOne(newTodo)
	usecase.outputPort.Write(newTodo, err)
}
