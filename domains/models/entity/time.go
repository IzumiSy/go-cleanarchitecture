package entity

import (
	"time"
)

type Time_ struct {
	// [時間系値オブジェクトの抽象]
	// これをラップしてCreatedAtとかModifiedAtとかつくる

	value time.Time
}

func NewTime(value time.Time) Time_ {
	return Time_{value: value}
}

func GenerateTime() Time_ {
	return Time_{value: time.Now()}
}

func (t Time_) Value() time.Time {
	return t.value
}
