package models

import (
	"go-cleanarchitecture/domains/models/session"
)

type Session struct {
	// ユーザーの認証セッションを表現するエンティティ

	id        session.ID
	user      User
	createdAt session.CreatedAt
}

func NewSession(user User) Session {
	return Session{
		id:        session.GenerateID(),
		user:      user,
		createdAt: session.GenerateCreatedAt(),
	}
}
