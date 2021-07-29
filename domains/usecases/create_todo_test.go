package usecases

import (
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/domains/models/todo"
	"go-cleanarchitecture/testing/adapters"
	"go-cleanarchitecture/testing/adapters/dao"
	"testing"
)

type mockCreateTodoOutputPort struct {
	Result errors.Domain
}

func (m mockCreateTodoOutputPort) Raise(err errors.Domain) {
	m.Result = err
}

func (_ mockCreateTodoOutputPort) Write(todo models.Todo) {}

func TestCreateTodoUsecase(t *testing.T) {
	p := &mockCreateTodoOutputPort{}

	usecase := CreateTodoUsecase{
		OutputPort: p,
		TodoDao: dao.MockTodoDao{
			GetResult: func() (models.Todo, errors.Domain, bool) {
				return models.Todo{}, errors.None, false
			},
			GetByNameResult: func() (models.Todo, errors.Domain, bool) {
				name, _ := todo.NewName("testing todo")
				description, _ := todo.NewDescription("this is a testing todo")
				todo := models.NewTodo(name, description)
				return todo, errors.None, true
			},
			StoreResult: func() errors.Domain {
				return errors.None
			},
		},
		TodosDao:  dao.MockTodosDao{},
		Logger:    adapters.MockLogger{T: t},
		Publisher: adapters.MockPublisher{},
	}

	usecase.
		Build(CreateTodoParam{Name: "testing todo", Description: "this is a testing todo"}).
		Run(adapters.MockAuthorizer{})

	if !p.Result.IsType(errors.PreconditionalError) {
		t.Errorf("Error must be preconditional (caught: %s)", p.Result.Error())
	}
}
