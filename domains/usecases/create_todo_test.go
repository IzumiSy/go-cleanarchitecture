package usecases

import (
	"go-cleanarchitecture/domains"
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/domains/models/todo"
	"go-cleanarchitecture/domains/models/user"
	"go-cleanarchitecture/testing/adapters"
	"go-cleanarchitecture/testing/adapters/dao"
	"testing"

	"golang.org/x/xerrors"
)

type mockCreateTodoOutputPort struct {
	Result errors.Domain
}

func (m *mockCreateTodoOutputPort) Raise(err errors.Domain) {
	m.Result = err
}

func (_ *mockCreateTodoOutputPort) Write(todo models.Todo) {}

func TestCreateTodoUsecase(t *testing.T) {
	newUsecase := func(
		todoDao domains.TodoRepository,
		todosDao domains.TodosRepository,
		op *mockCreateTodoOutputPort,
	) CreateTodoUsecase {
		return CreateTodoUsecase{
			OutputPort: op,
			TodoDao:    todoDao,
			TodosDao:   todosDao,
			Logger:     adapters.MockLogger{T: t},
			Publisher:  adapters.MockPublisher{},
		}
	}

	t.Run(uc_TODO_NAME_NOT_UNIQUE.Reason(), func(t *testing.T) {
		todoDao := dao.NewMockTodoDao()
		todoDao.GetByNameResult = func() (models.Todo, errors.Domain, bool) {
			name, _ := todo.NewName("testing todo")
			description, _ := todo.NewDescription("this is a testing todo")
			userID, _ := user.NewID("user_id")
			todo := models.NewTodo(name, description, userID)
			return todo, errors.None, true
		}

		op := &mockCreateTodoOutputPort{}
		newUsecase(todoDao, dao.NewMockTodosDao(), op).
			Build(CreateTodoParam{Name: "testing todo", Description: "this is a testing todo"}).
			Run(adapters.MockAuthorizer{})

		if !xerrors.Is(op.Result, uc_TODO_NAME_NOT_UNIQUE) {
			t.Errorf("Unexpected validation error (caught: %s)", op.Result.Error())
		}
	})

	t.Run(uc_MAXIMUM_TODOS_REACHED.Reason(), func(t *testing.T) {
		todosDao := dao.NewMockTodosDao()
		todosDao.GetByUserIDResult = func() (models.Todos, errors.Domain) {
			todos := make([]models.Todo, 100)
			return models.NewTodos(todos), errors.None
		}

		op := &mockCreateTodoOutputPort{}
		newUsecase(dao.NewMockTodoDao(), todosDao, op).
			Build(CreateTodoParam{Name: "testing todo", Description: "this is a testing todo"}).
			Run(adapters.MockAuthorizer{})

		if !xerrors.Is(op.Result, uc_MAXIMUM_TODOS_REACHED) {
			t.Errorf("Unexpected validation error (caught: %s)", op.Result.Error())
		}
	})
}
