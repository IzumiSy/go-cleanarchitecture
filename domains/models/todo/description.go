package todo

import (
	"go-cleanarchitecture/domains/errors"
	"unicode/utf8"
)

type Description struct {
	// [TODOの説明を表現する値オブジェクト]
	// バリデーションルールは以下
	// - 空ではないこと
	// - 100文字以下であること

	value string
}

func NewDescription(value string) (Description, errors.Domain) {
	empty := Description{}

	if value == "" {
		return empty, errors.Invalid("Description must not be empty")
	} else if utf8.RuneCountInString(value) > 100 {
		return empty, errors.Invalid("Description too long")
	} else {
		return Description{value}, errors.None
	}
}

func (description Description) Value() string {
	return description.value
}
