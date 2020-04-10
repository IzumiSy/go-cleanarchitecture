package json

import (
	"go-cleanarchitecture/adapters/presenters"
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models"
)

type createTodoResponse struct {
	Id string `json:"id"`
}

type CreateTodoPresenter struct {
	Presenter presenters.EchoPresenter
}

// エラーハンドリングはサボって全部500を返している
func (presenter CreateTodoPresenter) Write(todo models.Todo) {
	presenter.Presenter.Succeed(createTodoResponse{
		Id: todo.Id().String(),
	})
}

func (presenter CreateTodoPresenter) Raise(reason errors.Domain) {
	presenter.Presenter.Fail()
}

func (presenter CreateTodoPresenter) Present() error {
	return presenter.Presenter.Result()
}
