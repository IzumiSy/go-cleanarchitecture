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
	logger     domains.Logger
}

func NewGetTodosUsecase(
	outputPort GetTodosOutputPort,
	todosDao domains.TodosRepository,
	logger domains.Logger,
) getTodosUsecase {
	return getTodosUsecase{outputPort, todosDao, logger}
}

func (usecase getTodosUsecase) Execute() {
	todos, err, _ := usecase.todosDao.Get()
	if err.NotNil() {
		usecase.outputPort.Raise(err)
		return
	}

	usecase.outputPort.Write(todos)
}
