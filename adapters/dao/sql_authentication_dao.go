package dao

import (
	"go-cleanarchitecture/domains"
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/domains/models/authentication"
	"go-cleanarchitecture/domains/models/user"
	"time"
)

type AuthentcationDao SQLDao

var _ domains.AuthenticationRepository = AuthentcationDao{}

func NewSQLAuthenticationDao(tt txType) (AuthentcationDao, error) {
	dao, err := newSQLDao("authentication", tt)
	return AuthentcationDao(dao), err
}

func (dao AuthentcationDao) Close() {
	dao.Close()
}

type authenticationDto struct {
	Email     string    `gorm:"email,primaryKey,uniqueIndex"`
	Hash      string    `gorm:"hash"`
	UserID    string    `gorm:"user_id"`
	CreatedAt time.Time `gorm:"created_at"`
}

type userDto struct {
	ID   string `gorm:"id"`
	Name string `gorm:"name"`
}

func (dao AuthentcationDao) GetByEmail(email authentication.Email) (models.Authentication, errors.Domain, bool) {
	var authDto authenticationDto

	query := dao.
		conn.
		Where("email = ?", email.Value()).
		Find(&authDto)

	empty := models.Authentication{}

	if query.RecordNotFound() {
		return empty, errors.None, false
	} else if query.Error != nil {
		return empty, errors.External(query.Error), false
	}

	var userDto userDto

	query = dao.
		conn.
		Table("user").
		Where("id = ?", authDto.UserID).
		Find(&userDto)

	if query.RecordNotFound() {
		return empty, errors.None, false
	} else if query.Error != nil {
		return empty, errors.External(query.Error), false
	}

	// 永続化済みのデータの取り出しでバリデーションエラーはないはずなのでエラーは無視する
	_email, _ := authentication.NewEmail(authDto.Email)
	hash := authentication.NewHash(authDto.Hash)
	createdAt := authentication.NewCreatedAt(authDto.CreatedAt)
	userID, _ := user.NewID(userDto.ID)
	userName, _ := user.NewName(userDto.Name)
	user := models.BuildUser(userID, userName)

	return models.BuildAuthentication(_email, hash, user, createdAt), errors.None, true
}

func (dao AuthentcationDao) Store(auth models.Authentication) errors.Domain {
	user := auth.User()
	authDto := authenticationDto{
		Email:     auth.Email().Value(),
		Hash:      auth.Hash().Value(),
		UserID:    user.ID().String(),
		CreatedAt: auth.CreatedAt().Value(),
	}
	if err := dao.conn.Create(&authDto).Error; err != nil {
		return errors.External(err)
	}

	userDto := userDto{
		ID:   user.ID().String(),
		Name: user.Name().Value(),
	}
	if err := dao.conn.Table("user").Create(&userDto).Error; err != nil {
		return errors.External(err)
	}

	return errors.None
}
