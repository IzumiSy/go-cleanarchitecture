package session

import (
	"go-cleanarchitecture/domains/models/entity"
)

type ID struct {
	entity.ID_
}

type CreatedAt struct {
	// 認証セッションの作成日時を表す値オブジェクト
	entity.Time_
}
