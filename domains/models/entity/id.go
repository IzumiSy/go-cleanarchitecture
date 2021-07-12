package entity

import (
	"fmt"
	"go-cleanarchitecture/domains/errors"

	"github.com/google/uuid"
)

type ID struct {
	// [エンティティの識別子を表現する値オブジェクトの抽象]
	// いま時点ではUUID型をラップしているが今後IDの実装が変わった際でも
	// 変更範囲を個々だけの留めることができる。

	value uuid.UUID
}

func NewID(value string) (ID, errors.Domain) {
	id, err := uuid.Parse(value)
	if err != nil {
		return ID{}, errors.Invalid(fmt.Sprintf("Invalid User ID: %s", err.Error()))
	}

	return ID{id}, errors.None
}

func (id ID) String() string {
	return id.value.String()
}

// IDを生成する
func GenerateID() ID {
	return ID{uuid.New()}
}
