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

func (id ID_) String() string {
	return id.value.String()
}

type IDBuilder interface {
	Build() (ID_, errors.Domain)
}

var (
	_ IDBuilder = ParseID{}
	_ IDBuilder = GenerateID{}
)

type ParseID struct {
	Src string
}

func (v ParseID) Build() (ID_, errors.Domain) {
	if id, err := uuid.Parse(v.Src); err != nil {
		return ID_{}, errors.Preconditional(fmt.Sprintf("IDBuilder: invalid ID: %s", err.Error()))
	} else {
		return ID_{value: id}, errors.None
	}
}

type GenerateID struct{}

func (GenerateID) Build() (ID_, errors.Domain) {
	return ID_{value: uuid.New()}, errors.None
}
