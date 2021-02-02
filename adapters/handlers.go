package adapters

import (
	"github.com/labstack/echo"
	"go-cleanarchitecture/adapters/dao"
	"go-cleanarchitecture/adapters/loggers"
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

	logger, err := loggers.NewZapLogger("config/zap.json")
	if err != nil {
		return err
	}

	presenter := json.GetTodosPresenter{Presenter: presenters.NewPresenter(ctx)}
	usecases.NewGetTodosUsecase(presenter, sqlDao, logger).Execute()
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
	// defer sqlDao.Close()

	logger, err := loggers.NewZapLogger("config/zap.json")
	if err != nil {
		return err
	}

	presenter := json.CreateTodoPresenter{Presenter: presenters.NewPresenter(ctx)}
	usecases.
		NewCreateTodoUsecase(presenter, sqlDao, logger).
		Execute(usecases.CreateTodoParam{
			Name:        jsonParam.Name,
			Description: jsonParam.Description,
		})
	return presenter.Present()
}
