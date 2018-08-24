package adapters

import (
	"github.com/IzumiSy/go-cleanarchitecture/adapters/dao"
	"github.com/IzumiSy/go-cleanarchitecture/adapters/presenters"
	"github.com/IzumiSy/go-cleanarchitecture/domains/usecases"
	"github.com/labstack/echo"
)

func getTodosHandler(ctx echo.Context) error {
	jsonPresenter := presenters.NewGetTodosJsonPresenter(ctx)
	sqlTodoDao := dao.NewSqlTodoDao()
	usecase := usecases.NewGetTodosUsecase(jsonPresenter, sqlTodoDao)
	usecase.Execute()
	return jsonPresenter.Present()
}
