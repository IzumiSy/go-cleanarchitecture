package adapters

import (
	"github.com/labstack/echo"
	"go-cleanarchitecture/adapters/dao"
	"go-cleanarchitecture/adapters/presenters"
	"go-cleanarchitecture/adapters/presenters/json"
	"go-cleanarchitecture/domains/usecases"
)

func getTodosHandler(ctx echo.Context) error {
	sqlDao, err := dao.NewSQLTodosDao()
	if err != nil {
		return err
	}
	defer sqlDao.Close()

	presenter := json.GetTodosPresenter{presenters.NewPresenter(ctx)}
	usecases.NewGetTodosUsecase(presenter, sqlDao).Execute()
	return presenter.Present()
}

type jsonCreateTodoParam struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func createTodoHandler(ctx echo.Context) error {
	jsonParam := new(jsonCreateTodoParam)

	if err := ctx.Bind(jsonParam); err != nil {
		return err
	}

	sqlDao, err := dao.NewSQLTodoDao()
	if err != nil {
		return err
	}
	defer sqlDao.Close()

	presenter := json.CreateTodoPresenter{presenters.NewPresenter(ctx)}
	usecases.
		NewCreateTodoUsecase(presenter, sqlDao).
		Execute(usecases.CreateTodoParam{
			Name:        jsonParam.Name,
			Description: jsonParam.Description,
		})
	return presenter.Present()
}
