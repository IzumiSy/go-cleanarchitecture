package todo

import (
	"errors"
	"unicode/utf8"
)

type Description struct {
	// [TODOの説明を表現する値オブジェクト]
	// バリデーションロジックは以下
	// - 説明は空ではないこと
	// - 説明は100文字以下であること

	value string
}

func NewDescription(value string) (error, Description) {
	empty := Description{}

	if value == "" {
		return errors.New("Description must not be empty"), empty
	} else if utf8.RuneCountInString(value) > 100 {
		return errors.New("Description too long"), empty
	} else {
		return nil, Description{value}
	}
}

func (description Description) Value() string {
	return description.value
}
