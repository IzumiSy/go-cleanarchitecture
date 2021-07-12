package dao

import (
	"go-cleanarchitecture/domains"
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/domains/models/authentication"
)

type AuthentcationDao SQLDao

var _ domains.AuthenticationRepository = AuthentcationDao{}

func NewSQLAuthenticationDao(tt txType) (AuthentcationDao, error) {
	err, dao := newSQLDao(tt)
	return AuthentcationDao(dao), err
}

func (dao AuthentcationDao) Close() {
	dao.Close()
}

type authenticationDto struct {
	Email     string `gorm:"email"`
	Hash      string `gorm:"hash"`
	CreatedAt string `gorm:"created_at"`
}

func (dao AuthentcationDao) GetByEmail(email authentication.Email) (models.Authentication, errors.Domain, bool) {
	var auth models.Authentication

	// todo

	return auth, errors.None, true
}

func (dao AuthentcationDao) Store(auth models.Authentication) errors.Domain {
	// todo

	return errors.None
}
