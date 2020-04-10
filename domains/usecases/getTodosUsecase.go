package usecases

import (
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/domains/repositories"
)

type GetTodosOutputPort interface {
	OutputPort
	Write(todos []models.Todo, result error)
}

type getTodosUsecase struct {
	outputPort GetTodosOutputPort
	todoDao    repositories.TodoRepository
}

func NewGetTodosUsecase(outputPort GetTodosOutputPort, todoDao repositories.TodoRepository) getTodosUsecase {
	return getTodosUsecase{outputPort, todoDao}
}

func (usecase getTodosUsecase) Execute() {
	todos, err := usecase.todoDao.GetAll()
	usecase.outputPort.Write(todos, err)
}
