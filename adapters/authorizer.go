package adapters

import (
	"errors"
	"go-cleanarchitecture/adapters/dao"
	"go-cleanarchitecture/domains"
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/domains/models/entity"
	"go-cleanarchitecture/domains/models/session"
	"net/http"
	"strings"
)

type DBSessionAuthorizer struct {
	Request *http.Request
	Driver  dao.Driver
}

var _ domains.Authorizer = DBSessionAuthorizer{}

func (da DBSessionAuthorizer) Run() (models.Session, error) {
	sessionDao, err := da.Driver.NewSQLSessionDao(dao.WITHOUT_TX())
	if err != nil {
		return models.Session{}, err
	}
	// defer sessionDao.Close()

	token, err := da.extractBearerToken()
	if err != nil {
		return models.Session{}, err
	}

	sessionID, sErr := session.NewID(entity.ParseID{Src: token})
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
	bearerToken := da.Request.Header.Get("Authorization")
	token := strings.Split(bearerToken, "Bearer ")
	if len(token) == 2 {
		return token[1], nil
	}
	return "", errors.New("Invalid authorization token")
}
