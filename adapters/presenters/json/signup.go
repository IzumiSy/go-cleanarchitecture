package json

import (
	"go-cleanarchitecture/adapters/presenters"
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/domains/usecases"
)

type signupResponse struct {
	CreatedAt int64 `json:"created_at"`
}

type SignupPresenter struct {
	Presenter presenters.EchoPresenter
}

var _ usecases.SignupOutputPort = SignupPresenter{}

func (presenter SignupPresenter) Write(auth models.Authentication) {
	presenter.Presenter.Succeed(signupResponse{
		CreatedAt: auth.CreatedAt().Value().Unix(),
	})
}

func (presenter SignupPresenter) Raise(reason errors.Domain) {
	// エラーハンドリングはサボって全部500を返している
	presenter.Presenter.Fail()
}

func (presenter SignupPresenter) Present() error {
	return presenter.Presenter.Result()
}
