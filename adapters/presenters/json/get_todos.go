package json

import (
	"go-cleanarchitecture/adapters/presenters"
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/domains/usecases"
)

type getTodosResponse struct {
	Todos []getTodosResponseItem `json:"todos"`
}

type getTodosResponseItem struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GetTodosPresenter struct {
	Presenter presenters.EchoPresenter
}

var _ usecases.GetTodosOutputPort = GetTodosPresenter{}

func (presenter GetTodosPresenter) Write(todos models.Todos) {
	// nilではなく空配列でレスポンスを返せるようにする
	response := getTodosResponse{
		Todos: []getTodosResponseItem{},
	}

	for _, todo := range todos.Value() {
		response.Todos = append(response.Todos, getTodosResponseItem{
			ID:          todo.ID().String(),
			Name:        todo.Name().Value(),
			Description: todo.Description().Value(),
		})
	}

	presenter.Presenter.Succeed(response)
}

func (presenter GetTodosPresenter) Raise(reason errors.Domain) {
	presenter.Presenter.Fail(reason)
}
