package adapters

import (
	"go-cleanarchitecture/adapters/dao"
	"go-cleanarchitecture/adapters/presenters"
	"go-cleanarchitecture/domains/usecases"
	"github.com/labstack/echo"
)

func getTodosHandler(ctx echo.Context) error {
	jsonPresenter := presenters.NewGetTodosJsonPresenter(ctx)
	sqlTodoDao := dao.NewSqlTodoDao()
	usecase := usecases.NewGetTodosUsecase(jsonPresenter, sqlTodoDao)
	usecase.Execute()
	return jsonPresenter.Present()
}
