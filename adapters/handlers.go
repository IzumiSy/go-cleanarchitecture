package adapters

import (
	"go-cleanarchitecture/adapters/dao"
	"go-cleanarchitecture/adapters/loggers"
	"go-cleanarchitecture/adapters/presenters"
	"go-cleanarchitecture/adapters/presenters/json"
	"go-cleanarchitecture/domains"
	"go-cleanarchitecture/domains/usecases"

	"github.com/labstack/echo"
)

type Handler = func(ctx echo.Context) error

func signupHandler(publisher domains.EventPublisher) Handler {
	return func(ctx echo.Context) error {
		jsonParam := new(struct {
			Email    string `json:"email"`
			Password string `json:"password"`
			UserName string `json:"userName"`
		})

		if err := ctx.Bind(jsonParam); err != nil {
			return err
		}

		return dao.WithTx(func(tx dao.TxSQLDao) error {
			authenticationDao, err := dao.NewSQLAuthenticationDao(dao.WITH_TX(tx))
			if err != nil {
				return err
			}
			// defer authenticationDao.Close()

			logger, err := loggers.NewZapLogger("config/zap.json")
			if err != nil {
				return err
			}

			presenter := json.SignupPresenter{Presenter: presenters.NewPresenter(ctx)}
			usecases.SignupUsecase{
				OutputPort:        presenter,
				AuthenticationDao: authenticationDao,
				Logger:            logger,
				Publisher:         publisher,
			}.Build(usecases.SignupParam{
				Email:    jsonParam.Email,
				Password: jsonParam.Password,
				UserName: jsonParam.UserName,
			}).Run()

			return presenter.Presenter.Result()
		})
	}
}

func authenticateHandler(publisher domains.EventPublisher) Handler {
	return func(ctx echo.Context) error {
		jsonParam := new(struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		})

		if err := ctx.Bind(jsonParam); err != nil {
			return err
		}

		return dao.WithTx(func(tx dao.TxSQLDao) error {
			authenticationDao, err := dao.NewSQLAuthenticationDao(dao.WITH_TX(tx))
			if err != nil {
				return err
			}
			// defer authenticationDao.Close()

			sessionDao, err := dao.NewSQLSessionDao(dao.WITH_TX(tx))
			if err != nil {
				return err
			}
			// defer sessionDao.Close()

			logger, err := loggers.NewZapLogger("config/zap.json")
			if err != nil {
				return err
			}

			presenter := json.AuthenticatePresenter{Presenter: presenters.NewPresenter(ctx)}
			usecases.AuthenticateUsecase{
				OutputPort:        presenter,
				AuthenticationDao: authenticationDao,
				SessionDao:        sessionDao,
				Logger:            logger,
				Publisher:         publisher,
			}.Build(usecases.AuthenticateParam{
				Email:    jsonParam.Email,
				Password: jsonParam.Password,
			}).Run()

			return presenter.Presenter.Result()
		})
	}
}

func getTodosHandler(ctx echo.Context) error {
	sqlDao, err := dao.NewSQLTodosDao(dao.WITHOUT_TX())
	if err != nil {
		return err
	}
	// defer sqlDao.Close()

	logger, err := loggers.NewZapLogger("config/zap.json")
	if err != nil {
		return err
	}

	presenter := json.GetTodosPresenter{Presenter: presenters.NewPresenter(ctx)}
	usecases.GetTodosUsecase{
		OutputPort: presenter,
		TodosDao:   sqlDao,
		Logger:     logger,
	}.Build().Run(DBSessionAuthorizer{ctx})

	return presenter.Presenter.Result()
}

func createTodoHandler(publisher domains.EventPublisher) Handler {
	return func(ctx echo.Context) error {
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
			// defer sqlTodoDao.Close()

			sqlTodosDao, err := dao.NewSQLTodosDao(dao.WITH_TX(tx))
			if err != nil {
				return err
			}
			// defer sqlTodosDao.Close()

			logger, err := loggers.NewZapLogger("config/zap.json")
			if err != nil {
				return err
			}

			presenter := json.CreateTodoPresenter{Presenter: presenters.NewPresenter(ctx)}
			usecases.CreateTodoUsecase{
				OutputPort: presenter,
				TodoDao:    sqlTodoDao,
				TodosDao:   sqlTodosDao,
				Logger:     logger,
				Publisher:  publisher,
			}.Build(usecases.CreateTodoParam{
				Name:        jsonParam.Name,
				Description: jsonParam.Description,
			}).Run(DBSessionAuthorizer{ctx})

			return presenter.Presenter.Result()
		})
	}
}

func signedUpHandler(payload []byte) error {
	return nil
}

func userAuthenticatedHandler(payload []byte) error {
	return nil
}

func todoCreatedHandler(payload []byte) error {
	return nil
}
