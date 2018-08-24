package usecases

import (
	"github.com/IzumiSy/go-cleanarchitecture/dao"
	"github.com/IzumiSy/go-cleanarchitecture/domains/models"
)

type GetTodosOutputPort interface {
	Write(todos []models.Todo, err error)
}

type getTodosUsecase struct {
	outputPort GetTodosOutputPort
	todoDao    dao.TodoDao
}

func NewGetTodosUsecase(outputPort GetTodosOutputPort, todoDao dao.TodoDao) getTodosUsecase {
	return getTodosUsecase{outputPort, todoDao}
}

func (usecase getTodosUsecase) Execute() {
	todos, err := usecase.todoDao.GetAll()
	usecase.outputPort.Write(todos, err)
}
