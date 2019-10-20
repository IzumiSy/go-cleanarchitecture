package usecases

import (
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/domains/repositories"
)

type CreateTodoOutputPort interface {
	Write(result error)
}

type CreateTodoParam struct {
	Name string
}

type createTodoUsecase struct {
	params     CreateTodoParam
	outputPort CreateTodoOutputPort
	todoDao    repositories.TodoRepository
}

func NewCreateTodoUsecase(
	params CreateTodoParam, outputPort CreateTodoOutputPort, todoDao repositories.TodoRepository,
) createTodoUsecase {
	return createTodoUsecase{params, outputPort, todoDao}
}

func (usecase createTodoUsecase) Execute() {
	newTodo := models.NewTodo(usecase.params.Name)
	usecase.outputPort.Write(usecase.todoDao.CreateOne(newTodo))
}
