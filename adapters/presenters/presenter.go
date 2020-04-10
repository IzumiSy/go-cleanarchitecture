package presenters

import (
	"github.com/labstack/echo"
)

type Presenter interface {
	Present() error
}

type EchoPresenter struct {
	ctx    echo.Context
	result error
}

func NewPresenter(ctx echo.Context) EchoPresenter {
	return EchoPresenter{ctx, nil}
}

func (presenter EchoPresenter) Fail() {
	presenter.result = presenter.ctx.String(500, "Internal Server Error")
}

func (presenter EchoPresenter) Succeed(value interface{}) {
	presenter.result = presenter.ctx.JSON(200, value)
}

func (presenter EchoPresenter) Result() error {
	return presenter.result
}
