package json

import (
	"go-cleanarchitecture/adapters/presenters"
	"go-cleanarchitecture/domains/models"
)

type getTodosResponse struct {
	Todos []getTodosResponseItem `json:"todos"`
}

type getTodosResponseItem struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type GetTodosPresenter struct {
	Presenter presenters.EchoPresenter
}

// エラーハンドリングはサボって全部500を返している
func (presenter GetTodosPresenter) Write(todos models.Todos, err error) {
	if err != nil {
		presenter.Presenter.Fail()
		return
	}

	var response getTodosResponse

	for _, todo := range todos.Value() {
		response.Todos = append(response.Todos, getTodosResponseItem{
			Id:   todo.Id().String(),
			Name: todo.Name().Value(),
		})
	}

	presenter.Presenter.Succeed(response)
}

func (presenter GetTodosPresenter) Raise(reason error) {
	presenter.Presenter.Fail()
}

func (presenter GetTodosPresenter) Present() error {
	return presenter.Presenter.Result()
}
