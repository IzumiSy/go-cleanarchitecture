package json

import (
	"go-cleanarchitecture/adapters/presenters"
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/domains/usecases"
)

type authenticateResponse struct {
	ID        string `json:"session_id"`
	CreatedAt int64  `json:"created_at"`
}

type AuthenticatePresenter struct {
	Presenter presenters.EchoPresenter
}

var _ usecases.AuthenticateOutputPort = AuthenticatePresenter{}

func (presenter AuthenticatePresenter) Write(session models.Session) {
	presenter.Presenter.Succeed(authenticateResponse{
		ID:        session.ID().String(),
		CreatedAt: session.CreatedAt().Value().Unix(),
	})
}

func (presenter AuthenticatePresenter) Raise(reason errors.Domain) {
	// エラーハンドリングはサボって全部500を返している
	presenter.Presenter.Fail()
}

func (presenter AuthenticatePresenter) Present() error {
	return presenter.Presenter.Result()
}
