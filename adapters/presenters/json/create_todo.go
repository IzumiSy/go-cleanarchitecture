package json

import (
	"go-cleanarchitecture/adapters/presenters"
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/domains/usecases"
)

type createTodoResponse struct {
	ID string `json:"id"`
}

type CreateTodoPresenter struct {
	Presenter presenters.EchoPresenter
}

var _ usecases.CreateTodoOutputPort = CreateTodoPresenter{}

func (presenter CreateTodoPresenter) Write(todo models.Todo) {
	presenter.Presenter.Succeed(createTodoResponse{
		ID: todo.ID().String(),
	})
}

func (presenter CreateTodoPresenter) Raise(reason errors.Domain) {
	presenter.Presenter.Fail(reason)
}
