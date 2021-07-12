package session

import "time"

type CreatedAt struct {
	// 認証セッションの作成日時を表す値オブジェクト

	value time.Time
}

func NewCreatedAt(value time.Time) CreatedAt {
	return CreatedAt{value}
}

func GenerateCreatedAt() CreatedAt {
	return CreatedAt{time.Now()}
}

func (c CreatedAt) Value() time.Time {
	return c.value
}
