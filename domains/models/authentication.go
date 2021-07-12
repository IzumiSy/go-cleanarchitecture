package models

import (
	"go-cleanarchitecture/domains/models/authentication"
	"time"
)

type createdAt time.Time

type Authentication struct {
	// ユーザーの認証情報を表現するエンティティ
	// 認証はドメインではなく実装詳細のような気もしたが
	// 一旦利用者の動作のひとつとしてユースケースに表すことにした。

	email     authentication.Email
	hash      authentication.Hash
	createdAt createdAt
}

func NewAuthentication(email authentication.Email, hash authentication.Hash) Authentication {
	return Authentication{
		email:     email,
		hash:      hash,
		createdAt: createdAt(time.Now()),
	}
}
