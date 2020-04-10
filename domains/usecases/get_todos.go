package usecases

import (
	"go-cleanarchitecture/domains"
	"go-cleanarchitecture/domains/models"
)

type GetTodosOutputPort interface {
	domains.OutputPort
	Write(todos models.Todos)
}

type getTodosUsecase struct {
	outputPort GetTodosOutputPort
	todosDao   domains.TodosRepository
}

func NewGetTodosUsecase(outputPort GetTodosOutputPort, todosDao domains.TodosRepository) getTodosUsecase {
	return getTodosUsecase{outputPort, todosDao}
}

func (usecase getTodosUsecase) Execute() {
	todos, err := usecase.todosDao.Get()
	if err.NotNil() {
		usecase.outputPort.Raise(err)
		return
	}

	usecase.outputPort.Write(todos)
}
