package entity

import (
	"fmt"
	"github.com/google/uuid"
	"go-cleanarchitecture/domains/errors"
)

type ID_ struct {
	// [エンティティの識別子を表現する値オブジェクトの抽象]
	// いま時点ではUUID型をラップしているが今後IDの実装が変わった際でも
	// 変更範囲を個々だけの留めることができる。

	value uuid.UUID
}

func NewID(value string) (ID_, errors.Domain) {
	id, err := uuid.Parse(value)
	if err != nil {
		return ID_{}, errors.Preconditional(fmt.Sprintf("Invalid ID: %s", err.Error()))
	}

	return ID_{id}, errors.None
}

func (id ID_) String() string {
	return id.value.String()
}

func GenerateID() ID_ {
	return ID_{uuid.New()}
}
