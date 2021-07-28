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
	presenter.Presenter.Fail(reason)
}
