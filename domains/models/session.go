package models

import (
	"go-cleanarchitecture/domains/models/entity"
	"go-cleanarchitecture/domains/models/session"
	"go-cleanarchitecture/domains/models/user"
)

type Session struct {
	// ユーザーの認証セッションを表現するエンティティ

	id        session.ID
	userID    user.ID
	createdAt session.CreatedAt
}

func NewSession(user User) Session {
	id, _ := session.NewID(entity.GenerateID{})
	return Session{
		id:        id,
		userID:    user.ID(),
		createdAt: session.CreatedAt{Time_: entity.GenerateTime()},
	}
}

func BuildSession(
	id session.ID,
	userID user.ID,
	createdAt session.CreatedAt,
) Session {
	return Session{
		id:        id,
		userID:    userID,
		createdAt: createdAt,
	}
}

func (session Session) ID() session.ID {
	return session.id
}

func (session Session) UserID() user.ID {
	return session.userID
}

func (session Session) CreatedAt() session.CreatedAt {
	return session.createdAt
}
