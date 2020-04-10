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

type jsonCreateTodoParam struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func createTodoHandler(ctx echo.Context) error {
	jsonParam := new(jsonCreateTodoParam)

	if err := ctx.Bind(jsonParam); err != nil {
		return err
	}

	presenter := json.CreateTodoPresenter{presenters.NewPresenter(ctx)}
	usecases.
		NewCreateTodoUsecase(presenter, dao.NewSqlTodoDao()).
		Execute(usecases.CreateTodoParam{
			Name:        jsonParam.Name,
			Description: jsonParam.Description,
		})
	return presenter.Present()
}
