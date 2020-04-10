package adapters

import (
	"github.com/labstack/echo"
	"go-cleanarchitecture/adapters/dao"
	"go-cleanarchitecture/adapters/presenters"
	"go-cleanarchitecture/adapters/presenters/json"
	"go-cleanarchitecture/domains/usecases"
)

func getTodosHandler(ctx echo.Context) error {
	presenter := json.GetTodosPresenter{presenters.NewPresenter(ctx)}
	usecases.NewGetTodosUsecase(presenter, dao.NewSqlTodoDao()).Execute()
	return presenter.Present()
}

func createTodoHandler(ctx echo.Context) error {
	presenter := json.CreateTodoPresenter{presenters.NewPresenter(ctx)}
	param := usecases.CreateTodoParam{"aaa", "bbb"} // todo
	usecases.NewCreateTodoUsecase(presenter, dao.NewSqlTodoDao()).Execute(param)
	return presenter.Present()
}
