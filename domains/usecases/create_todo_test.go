package usecases

import (
	"context"
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
	Error  errors.Domain
	Result models.Todo
}

func (m *mockCreateTodoOutputPort) Raise(err errors.Domain) {
	m.Error = err
}

func (m *mockCreateTodoOutputPort) Write(todo models.Todo) {
	m.Result = todo
}

type mockCreateTodoPublisher struct {
	event TodoCreatedEvent
}

func (p *mockCreateTodoPublisher) Publish(event domains.Event) errors.Domain {
	p.event = event.(TodoCreatedEvent)
	return errors.None
}

func TestCreateTodoUsecase(t *testing.T) {
	newUsecase := func(
		todoDao domains.TodoRepository,
		todosDao domains.TodosRepository,
		op *mockCreateTodoOutputPort,
		p *mockCreateTodoPublisher,
	) CreateTodoUsecase {
		return CreateTodoUsecase{
			Ctx:        context.Background(),
			OutputPort: op,
			TodoDao:    todoDao,
			TodosDao:   todosDao,
			Logger:     adapters.MockLogger{T: t},
			Publisher:  p,
		}
	}

	t.Run("Validation", func(t *testing.T) {
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
			p := &mockCreateTodoPublisher{}
			newUsecase(todoDao, dao.NewMockTodosDao(), op, p).
				Build(CreateTodoParam{Name: "testing todo", Description: "this is a testing todo"}).
				Run(adapters.MockAuthorizer{})

			if !xerrors.Is(op.Error, uc_TODO_NAME_NOT_UNIQUE) {
				t.Errorf("Unexpected validation error (caught: %s)", op.Error)
			}

			if p.event != (TodoCreatedEvent{}) {
				t.Error("Error: event published")
			}
		})

		t.Run(uc_MAXIMUM_TODOS_REACHED.Reason(), func(t *testing.T) {
			todosDao := dao.NewMockTodosDao()
			todosDao.GetByUserIDResult = func() (models.Todos, errors.Domain) {
				todos := make([]models.Todo, 100)
				return models.NewTodos(todos), errors.None
			}

			op := &mockCreateTodoOutputPort{}
			p := &mockCreateTodoPublisher{}
			newUsecase(dao.NewMockTodoDao(), todosDao, op, p).
				Build(CreateTodoParam{Name: "testing todo", Description: "this is a testing todo"}).
				Run(adapters.MockAuthorizer{})

			if !xerrors.Is(op.Error, uc_MAXIMUM_TODOS_REACHED) {
				t.Errorf("Unexpected validation error (caught: %s)", op.Error)
			}

			if p.event != (TodoCreatedEvent{}) {
				t.Error("Error: event published")
			}
		})
	})

	t.Run("Store", func(t *testing.T) {
		name, _ := todo.NewName("testing todo")
		description, _ := todo.NewDescription("this is a testing todo")

		op := &mockCreateTodoOutputPort{}
		p := &mockCreateTodoPublisher{}
		newUsecase(dao.NewMockTodoDao(), dao.NewMockTodosDao(), op, p).
			Build(CreateTodoParam{Name: name.Value(), Description: description.Value()}).
			Run(adapters.MockAuthorizer{})

		if !xerrors.Is(op.Error, errors.None) {
			t.Errorf("Unexpected error raised: %s", op.Error)
		}

		if op.Result.Name() != name || op.Result.Description() != description {
			t.Error("Error: invalid result")
		}

		if p.event == (TodoCreatedEvent{}) {
			t.Error("Error: event not published")
		}
	})
}
