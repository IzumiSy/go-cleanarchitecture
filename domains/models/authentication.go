package models

import (
	"go-cleanarchitecture/domains/models/authentication"
	"go-cleanarchitecture/domains/models/entity"
	"go-cleanarchitecture/domains/models/user"
)

type Authentication struct {
	// ユーザーの認証情報を表現するエンティティ
	// 認証はドメインではなく実装詳細のような気もしたが
	// 一旦利用者の動作のひとつとしてユースケースに表すことにした。

	email     authentication.Email
	hash      authentication.Hash
	createdAt authentication.CreatedAt

	// 認証情報には必ずユーザーが紐づくため
	// Authentication集約としてUserを持つようにモデリングすることにした
	user User
}

// ファクトリ関数としてUserエンティティも同時に生成する
func NewAuthentication(
	email authentication.Email,
	hash authentication.Hash,
	userName user.Name,
) Authentication {
	return Authentication{
		email:     email,
		hash:      hash,
		user:      NewUser(userName),
		createdAt: authentication.CreatedAt{Time_: entity.GenerateTime()},
	}
}

func BuildAuthentication(
	email authentication.Email,
	hash authentication.Hash,
	user User,
	createdAt authentication.CreatedAt,
) Authentication {
	return Authentication{
		email:     email,
		hash:      hash,
		user:      user,
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

func (auth Authentication) User() User {
	return auth.user
}
