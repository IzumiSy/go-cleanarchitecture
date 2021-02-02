package usecases

import (
	"go-cleanarchitecture/testing/adapters/dao"
	"go-cleanarchitecture/testing/adapters/loggers"
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models"
	"testing"
)

type mockCreateTodoOutputPort struct{}

func (_ mockCreateTodoOutputPort) Raise(err errors.Domain) {}

func (_ mockCreateTodoOutputPort) Write(todo models.Todo) {}

func TestCreateTodoUseCase(t *testing.T) {
	usecase := NewCreateTodoUsecase(
		mockCreateTodoOutputPort{},
		dao.MockTodoDao{},
		loggers.MockLogger{},
	)

	param := CreateTodoParam{Name: "todo0", Description: "todo0"}
	usecase.Execute(param)
}
