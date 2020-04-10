package usecases

import (
	"go-cleanarchitecture/domains"
	"go-cleanarchitecture/domains/models"
)

type GetTodosOutputPort interface {
	OutputPort
	Write(todos models.Todos, result error)
}

type getTodosUsecase struct {
	outputPort GetTodosOutputPort
	todosDao   domains.TodosRepository
}

func NewGetTodosUsecase(outputPort GetTodosOutputPort, todosDao domains.TodosRepository) getTodosUsecase {
	return getTodosUsecase{outputPort, todosDao}
}

func (usecase getTodosUsecase) Execute() {
	err, todos := usecase.todosDao.Get()
	usecase.outputPort.Write(todos, err)
}
