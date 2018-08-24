package usecases

import (
	"github.com/IzumiSy/go-cleanarchitecture/dao"
	"github.com/IzumiSy/go-cleanarchitecture/domains/models"
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
	todoDao    dao.TodoDao
}

func NewCreateTodoUsecase(
	params CreateTodoParam, outputPort CreateTodoOutputPort, todoDao dao.TodoDao,
) createTodoUsecase {
	return createTodoUsecase{params, outputPort, todoDao}
}

func (usecase createTodoUsecase) Execute() {
	newTodo := models.NewTodo(usecase.params.Name)
	usecase.outputPort.Write(usecase.todoDao.CreateOne(newTodo))
}
