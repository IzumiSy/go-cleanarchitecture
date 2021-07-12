package user

import (
	"go-cleanarchitecture/domains/errors"
	"unicode/utf8"
)

type Name struct {
	// Userの名前を表現する値オブジェクト
	// バリデーションルールは以下
	// - 空ではないこと
	// - 20文字以下であること

	value string
}

func NewName(value string) (errors.Domain, Name) {
	empty := Name{}

	if value == "" {
		return errors.Invalid("Name must not be empty"), empty
	} else if utf8.RuneCountInString(value) > 20 {
		return errors.Invalid("Name too long"), empty
	} else {
		return errors.None, Name{value}
	}
}

func (name Name) Value() string {
	return name.value
}
