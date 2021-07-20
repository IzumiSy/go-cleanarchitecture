package dao

import (
	"context"
	"go-cleanarchitecture/domains"
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/domains/models/authentication"
	"go-cleanarchitecture/domains/models/user"
	"time"

	"gorm.io/gorm"
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

type AuthenticationDto struct {
	Email     string    `gorm:"email,primaryKey,uniqueIndex"`
	Hash      string    `gorm:"hash"`
	UserID    string    `gorm:"user_id"`
	CreatedAt time.Time `gorm:"created_at"`
}

func (AuthenticationDto) TableName() string {
	return "authentication"
}

type UserDto struct {
	ID   string `gorm:"id"`
	Name string `gorm:"name"`
}

func (UserDto) TableName() string {
	return "user"
}

func (dao AuthentcationDao) GetByEmail(email authentication.Email) (models.Authentication, errors.Domain, bool) {
	var authDto AuthenticationDto

	query := dao.
		conn.
		WithContext(context.Background()).
		Where("email = ?", email.Value()).
		Take(&authDto)

	empty := models.Authentication{}

	if query.Error == gorm.ErrRecordNotFound {
		return empty, errors.None, false
	} else if query.Error != nil {
		return empty, errors.External(query.Error), false
	}

	var userDto UserDto

	query = dao.
		conn.
		WithContext(context.Background()).
		Table("user").
		Where("id = ?", authDto.UserID).
		Take(&userDto)

	if query.Error == gorm.ErrRecordNotFound {
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
	authDto := AuthenticationDto{
		Email:     auth.Email().Value(),
		Hash:      auth.Hash().Value(),
		UserID:    user.ID().String(),
		CreatedAt: auth.CreatedAt().Value(),
	}

	if err := dao.conn.WithContext(context.Background()).Create(&authDto).Error; err != nil {
		return errors.External(err)
	}

	userDto := UserDto{
		ID:   user.ID().String(),
		Name: user.Name().Value(),
	}
	if err := dao.conn.WithContext(context.Background()).Table("user").Create(&userDto).Error; err != nil {
		return errors.External(err)
	}

	return errors.None
}
