package adapters

import (
	"go-cleanarchitecture/adapters/dao"
	"go-cleanarchitecture/domains"
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/domains/models/session"

	"errors"
	"strings"

	"github.com/labstack/echo"
)

type DBSessionAuthorizer struct {
	Ctx echo.Context
}

var _ domains.Authorizer = DBSessionAuthorizer{}

func (da DBSessionAuthorizer) Run() (models.Session, error) {
	sessionDao, err := dao.NewSQLSessionDao(dao.WITHOUT_TX())
	if err != nil {
		return models.Session{}, err
	}
	// defer sessionDao.Close()

	token, err := da.extractBearerToken()
	if err != nil {
		return models.Session{}, err
	}

	sessionID, sErr := session.NewID(token)
	if sErr.NotNil() {
		return models.Session{}, sErr
	}

	session, sErr, exists := sessionDao.Get(sessionID)
	if sErr.NotNil() {
		return models.Session{}, sErr
	}

	if !exists {
		return models.Session{}, errors.New("Invalid session")
	}

	return session, nil
}

func (da DBSessionAuthorizer) extractBearerToken() (string, error) {
	bearerToken := da.Ctx.Request().Header.Get("Authorization")
	token := strings.Split(bearerToken, "Bearer ")
	if len(token) == 2 {
		return token[1], nil
	}
	return "", errors.New("Invalid authorization token")
}
