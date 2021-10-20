package session

import (
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models/entity"
)

type ID struct {
	entity.ID_
}

func NewID(builder entity.IDBuilder) (ID, errors.Domain) {
	id, err := builder.Build()
	return ID{ID_: id}, err
}

type CreatedAt struct {
	// 認証セッションの作成日時を表す値オブジェクト
	entity.Time_
}
