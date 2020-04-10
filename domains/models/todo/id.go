package todo

import (
	"github.com/google/uuid"
	"go-cleanarchitecture/domains/errors"
)

type Id struct {
	// [TODOの識別子を表現する値オブジェクト]
	// いま時点ではUUID型をラップしているが今後IDの実装が変わった際でも
	// 変更範囲を個々だけの留めることができる。

	value uuid.UUID
}

func NewId(value string) (Id, errors.Domain) {
	id, err := uuid.FromBytes([]byte(value))
	if err != nil {
		return Id{}, errors.Invalid("Invalid value")
	}

	return Id{id}, errors.None
}

// IDを内部の依存した実装のデータとして取得する
// 文字列として欲しい場合はStringメソッドを使うこと
func (id Id) Value() uuid.UUID {
	return id.value
}

func (id Id) String() string {
	return id.value.String()
}

// IDを生成する
func GenerateId() Id {
	return Id{uuid.New()}
}
