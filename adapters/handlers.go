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
	sqlDao, err := dao.NewSQLTodosDao(dao.WITHOUT_TX())
	if err != nil {
		return err
	}
	defer sqlDao.Close()

	logger, err := loggers.NewZapLogger("config/zap.json")
	if err != nil {
		return err
	}

	presenter := json.GetTodosPresenter{Presenter: presenters.NewPresenter(ctx)}
	usecases.
		NewGetTodosUsecase(presenter, sqlDao, logger).
		Execute(usecases.GetTodosParam{
			UserID: "user_id",
		})
	return presenter.Present()
}

func createTodoHandler(ctx echo.Context) error {
	jsonParam := new(struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	})

	if err := ctx.Bind(jsonParam); err != nil {
		return err
	}

	return dao.WithTx(func(tx dao.TxSQLDao) error {
		sqlTodoDao, err := dao.NewSQLTodoDao(dao.WITH_TX(tx))
		if err != nil {
			return err
		}
		defer sqlTodoDao.Close()

		sqlTodosDao, err := dao.NewSQLTodosDao(dao.WITH_TX(tx))
		if err != nil {
			return err
		}
		defer sqlTodosDao.Close()

		logger, err := loggers.NewZapLogger("config/zap.json")
		if err != nil {
			return err
		}

		presenter := json.CreateTodoPresenter{Presenter: presenters.NewPresenter(ctx)}
		usecases.
			NewCreateTodoUsecase(presenter, sqlTodoDao, sqlTodosDao, logger).
			Execute(usecases.CreateTodoParam{
				Name:        jsonParam.Name,
				Description: jsonParam.Description,
			})

		return presenter.Present()
	})
}
