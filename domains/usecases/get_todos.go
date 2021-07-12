package usecases

import (
	"go-cleanarchitecture/domains"
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/domains/models/user"
)

type GetTodosOutputPort interface {
	domains.OutputPort
	Write(todos models.Todos)
}

type GetTodosParam struct {
	UserID string
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

func (usecase getTodosUsecase) Execute(params GetTodosParam) {
	userID, _ := user.New(params.UserID)

	todos, err := usecase.todosDao.GetByUserID(userID)
	if err.NotNil() {
		usecase.outputPort.Raise(err)
		return
	}

	usecase.outputPort.Write(todos)
}
