package dao

import (
	"context"
	"go-cleanarchitecture/domains"
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/domains/models/entity"
	"go-cleanarchitecture/domains/models/session"
	"go-cleanarchitecture/domains/models/user"
	"time"

	"gorm.io/gorm"
)

type SessionDao SQLDao

var _ domains.SessionRepository = SessionDao{}

func (driver Driver) NewSQLSessionDao(tt txType) (SessionDao, error) {
	dao, err := driver.newSQLDao("session", tt)
	return SessionDao(dao), err
}

func (dao SessionDao) Close() {
	dao.Close()
}

type SessionDto struct {
	ID        string    `gorm:"column:id"`
	UserID    string    `gorm:"column:user_id;not null;index"`
	CreatedAt time.Time `gorm:"column:created_at;not null"`
}

func (SessionDto) TableName() string {
	return "session"
}

func (dao SessionDao) Get(id session.ID) (models.Session, errors.Domain, bool) {
	var sessionDto SessionDto

	query := dao.
		conn.
		WithContext(context.Background()).
		Take(&sessionDto, "id = ?", id.String())

	empty := models.Session{}

	if query.Error == gorm.ErrRecordNotFound {
		return empty, errors.None, false
	} else if query.Error != nil {
		return empty, errors.Postconditional(query.Error), false
	}

	// 永続化済みのデータの取り出しでバリデーションエラーはないはずなのでエラーは無視する
	sessionID, _ := session.NewID(entity.ParseID{Src: sessionDto.ID})
	userID, _ := user.NewID(entity.ParseID{Src: sessionDto.UserID})
	createdAt := session.CreatedAt{Time_: entity.NewTime(sessionDto.CreatedAt)}
	return models.BuildSession(sessionID, userID, createdAt), errors.None, true
}

func (dao SessionDao) Store(session models.Session) errors.Domain {
	dto := SessionDto{
		ID:        session.ID().String(),
		UserID:    session.UserID().String(),
		CreatedAt: session.CreatedAt().Value(),
	}

	return errors.Postconditional(
		dao.
			conn.
			WithContext(context.Background()).
			Create(&dto).
			Error,
	)
}
