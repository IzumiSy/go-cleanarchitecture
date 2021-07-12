package dao

import (
	"go-cleanarchitecture/domains"
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/domains/models/authentication"
	"time"
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
	Email     string    `gorm:"email,primaryKey,uniqueIndex"`
	Hash      string    `gorm:"hash"`
	CreatedAt time.Time `gorm:"created_at"`
}

func (dao AuthentcationDao) GetByEmail(email authentication.Email) (models.Authentication, errors.Domain, bool) {
	var dto authenticationDto

	query := dao.
		conn.
		Table("authentication").
		Where("email = ?", email.Value()).
		Find(&dto)

	empty := models.Authentication{}

	if query.RecordNotFound() {
		return empty, errors.None, false
	} else if query.Error != nil {
		return empty, errors.External(query.Error), false
	}

	// 永続化済みのデータの取り出しでバリデーションエラーはないはずなので無視する
	_email, _ := authentication.NewEmail(dto.Email)
	hash := authentication.NewHash(dto.Hash)
	createdAt := authentication.NewCreatedAt(dto.CreatedAt)

	return models.BuildAuthentication(_email, hash, createdAt), errors.None, true
}

func (dao AuthentcationDao) Store(auth models.Authentication) errors.Domain {
	dto := authenticationDto{
		Email:     auth.Email().Value(),
		Hash:      auth.Hash().Value(),
		CreatedAt: auth.CreatedAt().Value(),
	}

	return errors.External(dao.conn.Table("authentication").Create(&dto).Error)
}
