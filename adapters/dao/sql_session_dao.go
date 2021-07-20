package dao

import (
	"go-cleanarchitecture/domains"
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/domains/models/session"
	"go-cleanarchitecture/domains/models/user"
	"time"

	"gorm.io/gorm"
)

type SessionDao SQLDao

var _ domains.SessionRepository = SessionDao{}

func NewSQLSessionDao(tt txType) (SessionDao, error) {
	dao, err := newSQLDao("session", tt)
	return SessionDao(dao), err
}

func (dao SessionDao) Close() {
	dao.Close()
}

type SessionDto struct {
	gorm.Model

	ID        string    `gorm:"id"`
	UserID    string    `gorm:"user_id"`
	CreatedAt time.Time `gorm:"created_at"`
}

func (dao SessionDao) Get(id session.ID) (models.Session, errors.Domain, bool) {
	var sessionDto SessionDto

	query := dao.
		conn.
		First(&sessionDto, "id = ?", id.String())

	empty := models.Session{}

	if query.Error == gorm.ErrRecordNotFound {
		return empty, errors.None, false
	} else if query.Error != nil {
		return empty, errors.External(query.Error), false
	}

	// 永続化済みのデータの取り出しでバリデーションエラーはないはずなのでエラーは無視する
	_id, _ := session.NewID(sessionDto.ID)
	userID, _ := user.NewID(sessionDto.UserID)
	createdAt := session.NewCreatedAt(sessionDto.CreatedAt)

	return models.BuildSession(_id, userID, createdAt), errors.None, true
}

func (dao SessionDao) Store(session models.Session) errors.Domain {
	dto := SessionDto{
		ID:        session.ID().String(),
		UserID:    session.UserID().String(),
		CreatedAt: session.CreatedAt().Value(),
	}
	if err := dao.conn.Create(&dto).Error; err != nil {
		return errors.External(err)
	}

	return errors.None
}
