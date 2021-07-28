package todo

import (
	"go-cleanarchitecture/domains/errors"
	"unicode/utf8"
)

type Name struct {
	// [TODOの名前を表現する値オブジェクト]
	// バリデーションルールは以下
	// - 空ではないこと
	// - 30文字以下であること

	value string
}

func NewName(value string) (Name, errors.Domain) {
	empty := Name{}

	if value == "" {
		return empty, errors.Preconditional("Name must not be empty")
	} else if utf8.RuneCountInString(value) > 30 {
		return empty, errors.Preconditional("Name too long")
	} else {
		return Name{value}, errors.None
	}
}

func (name Name) Value() string {
	return name.value
}
