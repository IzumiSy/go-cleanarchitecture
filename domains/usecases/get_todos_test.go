package usecases

import (
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/testing/adapters"
	"go-cleanarchitecture/testing/adapters/dao"
	"go-cleanarchitecture/testing/adapters/loggers"
	"testing"
)

type mockGetTodosOutputPort struct{}

func (_ mockGetTodosOutputPort) Raise(err errors.Domain) {}

func (_ mockGetTodosOutputPort) Write(todo models.Todos) {}

func TestGetTodosUsecase(t *testing.T) {
	usecase := GetTodosUsecase{
		OutputPort: mockGetTodosOutputPort{},
		TodosDao:   dao.MockTodosDao{},
		Logger:     loggers.MockLogger{},
	}

	usecase.
		Build().
		Run(adapters.MockAuthorizer{})
}
