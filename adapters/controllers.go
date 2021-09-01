package adapters

import (
	"github.com/labstack/echo/v4"
	"go-cleanarchitecture/adapters/dao"
	"go-cleanarchitecture/adapters/presenters"
	"go-cleanarchitecture/adapters/presenters/json"
	"go-cleanarchitecture/domains"
	"go-cleanarchitecture/domains/usecases"
)

type Controller = func(ctx echo.Context) error

func signupController(publisher domains.EventPublisher, logger domains.Logger) Controller {
	return func(e echo.Context) error {
		jsonParam := new(struct {
			Email    string `json:"email"`
			Password string `json:"password"`
			UserName string `json:"userName"`
		})

		if err := e.Bind(jsonParam); err != nil {
			return err
		}

		return dao.WithTx(func(tx dao.TxSQLDao) error {
			authenticationDao, err := dao.NewSQLAuthenticationDao(dao.WITH_TX(tx))
			if err != nil {
				return err
			}
			// defer authenticationDao.Close()

			presenter := json.SignupPresenter{Presenter: presenters.NewPresenter(e)}
			usecases.SignupUsecase{
				Ctx:               e.Request().Context(),
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

func authenticateController(publisher domains.EventPublisher, logger domains.Logger) Controller {
	return func(e echo.Context) error {
		jsonParam := new(struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		})

		if err := e.Bind(jsonParam); err != nil {
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

			presenter := json.AuthenticatePresenter{Presenter: presenters.NewPresenter(e)}
			usecases.AuthenticateUsecase{
				Ctx:               e.Request().Context(),
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

func getTodosController(logger domains.Logger) Controller {
	return func(e echo.Context) error {
		sqlDao, err := dao.NewSQLTodosDao(dao.WITHOUT_TX())
		if err != nil {
			return err
		}
		// defer sqlDao.Close()

		presenter := json.GetTodosPresenter{Presenter: presenters.NewPresenter(e)}
		usecases.GetTodosUsecase{
			Ctx:        e.Request().Context(),
			OutputPort: presenter,
			TodosDao:   sqlDao,
			Logger:     logger,
		}.Build().Run(DBSessionAuthorizer{e})

		return presenter.Presenter.Result()
	}
}

func createTodoController(publisher domains.EventPublisher, logger domains.Logger) Controller {
	return func(e echo.Context) error {
		jsonParam := new(struct {
			Name        string `json:"name"`
			Description string `json:"description"`
		})

		if err := e.Bind(jsonParam); err != nil {
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

			presenter := json.CreateTodoPresenter{Presenter: presenters.NewPresenter(e)}
			usecases.CreateTodoUsecase{
				Ctx:        e.Request().Context(),
				OutputPort: presenter,
				TodoDao:    sqlTodoDao,
				TodosDao:   sqlTodosDao,
				Logger:     logger,
				Publisher:  publisher,
			}.Build(usecases.CreateTodoParam{
				Name:        jsonParam.Name,
				Description: jsonParam.Description,
			}).Run(DBSessionAuthorizer{e})

			return presenter.Presenter.Result()
		})
	}
}
