package json

import (
	"go-cleanarchitecture/adapters/presenters"
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/domains/usecases"
)

type signupResponse struct{}

type SignupPresenter struct {
	Presenter presenters.EchoPresenter
}

var _ usecases.SignupOutputPort = SignupPresenter{}

func (presenter SignupPresenter) Write(auth models.Authentication) {
	presenter.Presenter.Succeed(signupResponse{})
}

func (presenter SignupPresenter) Raise(reason errors.Domain) {
	// エラーハンドリングはサボって全部500を返している
	presenter.Presenter.Fail()
}

func (presenter SignupPresenter) Present() error {
	return presenter.Presenter.Result()
}
