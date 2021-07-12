package models

import (
	"go-cleanarchitecture/domains/models/authentication"
)

type Authentication struct {
	// ユーザーの認証情報を表現するエンティティ
	// 認証はドメインではなく実装詳細のような気もしたが
	// 一旦利用者の動作のひとつとしてユースケースに表すことにした。

	email     authentication.Email
	hash      authentication.Hash
	createdAt authentication.CreatedAt
}

func NewAuthentication(email authentication.Email, hash authentication.Hash) Authentication {
	return Authentication{
		email:     email,
		hash:      hash,
		createdAt: authentication.GenerateCreatedAt(),
	}
}

func BuildAuthentication(
	email authentication.Email,
	hash authentication.Hash,
	createdAt authentication.CreatedAt,
) Authentication {
	return Authentication{
		email:     email,
		hash:      hash,
		createdAt: createdAt,
	}
}

func (auth Authentication) Email() authentication.Email {
	return auth.email
}

func (auth Authentication) Hash() authentication.Hash {
	return auth.hash
}

func (auth Authentication) CreatedAt() authentication.CreatedAt {
	return auth.createdAt
}
