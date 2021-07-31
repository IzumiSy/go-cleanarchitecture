package presenters

import (
	"go-cleanarchitecture/domains/errors"
	"net/http"

	"github.com/labstack/echo"
)

type Presenter interface {
	Result() error
}

type EchoPresenter struct {
	ctx echo.Context
	err error
}

func NewPresenter(ctx echo.Context) EchoPresenter {
	return EchoPresenter{ctx, nil}
}

func (presenter EchoPresenter) Fail(err errors.Domain) {
	presenter.err = err

	if err.IsType(errors.PreconditionalError) {
		presenter.ctx.String(http.StatusBadRequest, err.Reason())
	} else {
		presenter.ctx.String(http.StatusInternalServerError, err.Reason())
	}
}

func (presenter EchoPresenter) Succeed(value interface{}) {
	presenter.ctx.JSON(http.StatusOK, value)
}

func (presenter EchoPresenter) Result() error {
	return presenter.err
}
