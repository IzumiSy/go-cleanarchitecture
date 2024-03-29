package user

import (
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models/entity"
	"unicode/utf8"
)

// [UserエンティティのID]
type ID struct {
	entity.ID_
}

func NewID(builder entity.IDBuilder) (ID, errors.Domain) {
	id, err := builder.Build()
	return ID{ID_: id}, err
}

type Name struct {
	// Userの名前を表現する値オブジェクト
	// バリデーションルールは以下
	// - 空ではないこと
	// - 20文字以下であること

	value string
}

func NewName(value string) (Name, errors.Domain) {
	empty := Name{}

	if value == "" {
		return empty, errors.Preconditional("Name must not be empty")
	} else if utf8.RuneCountInString(value) > 20 {
		return empty, errors.Preconditional("Name too long")
	} else {
		return Name{value}, errors.None
	}
}

func (name Name) Value() string {
	return name.value
}
