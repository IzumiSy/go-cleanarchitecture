package adapters

import (
	"errors"
	"go-cleanarchitecture/adapters/dao"
	"go-cleanarchitecture/adapters/loggers"
	"go-cleanarchitecture/adapters/presenters"
	"go-cleanarchitecture/adapters/presenters/json"
	"go-cleanarchitecture/domains/usecases"
	"strings"

	"github.com/labstack/echo"
)

func extractBearerToken(ctx echo.Context) (string, error) {
	bearerToken := ctx.Request().Header.Get("Authorization")
	token := strings.Split(bearerToken, "Bearer ")
	if len(token) == 2 {
		return token[1], nil
	}
	return "", errors.New("Invalid authorization token")
}

func signupHandler(ctx echo.Context) error {
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
		}.Build(usecases.SignupParam{
			Email:    jsonParam.Email,
			Password: jsonParam.Password,
			UserName: jsonParam.UserName,
		}).Run()
		return presenter.Present()
	})
}

func authenticateHandler(ctx echo.Context) error {
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
		}.Build(usecases.AuthenticateParam{
			Email:    jsonParam.Email,
			Password: jsonParam.Password,
		}).Run()
		return presenter.Present()
	})
}

func getTodosHandler(ctx echo.Context) error {
	sqlDao, err := dao.NewSQLTodosDao(dao.WITHOUT_TX())
	if err != nil {
		return err
	}
	// defer sqlDao.Close()

	sessionDao, err := dao.NewSQLSessionDao(dao.WITHOUT_TX())
	if err != nil {
		return err
	}
	// defer sessionDao.Close()

	logger, err := loggers.NewZapLogger("config/zap.json")
	if err != nil {
		return err
	}

	presenter := json.GetTodosPresenter{Presenter: presenters.NewPresenter(ctx)}
	usecases.GetTodosUsecase{
		OutputPort: presenter,
		TodosDao:   sqlDao,
		Logger:     logger,
	}.Build().
		Run(sessionDao, "d70f4845-b645-4271-bea9-3d5705e79e87")

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
		// defer sqlTodoDao.Close()

		sqlTodosDao, err := dao.NewSQLTodosDao(dao.WITH_TX(tx))
		if err != nil {
			return err
		}
		// defer sqlTodosDao.Close()

		sessionDao, err := dao.NewSQLSessionDao(dao.WITHOUT_TX())
		if err != nil {
			return err
		}
		// defer sessionDao.Close()

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
		}.Build(usecases.CreateTodoParam{
			Name:        jsonParam.Name,
			Description: jsonParam.Description,
		}).Run(sessionDao, "d70f4845-b645-4271-bea9-3d5705e79e87")

		return presenter.Present()
	})
}
