package usecases

import (
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/testing/adapters"
	"go-cleanarchitecture/testing/adapters/dao"
	"go-cleanarchitecture/testing/adapters/loggers"
	"testing"
)

type mockCreateTodoOutputPort struct{}

func (_ mockCreateTodoOutputPort) Raise(err errors.Domain) {}

func (_ mockCreateTodoOutputPort) Write(todo models.Todo) {}

func TestCreateTodoUsecase(t *testing.T) {
	usecase := CreateTodoUsecase{
		OutputPort: mockCreateTodoOutputPort{},
		TodoDao:    dao.MockTodoDao{},
		TodosDao:   dao.MockTodosDao{},
		Logger:     loggers.MockLogger{},
	}

	usecase.
		Build(CreateTodoParam{Name: "todo0", Description: "todo0"}).
		Run(adapters.MockAuthorizer{})
}
