package todo

import (
	"errors"
	"unicode/utf8"
)

type Name struct {
	// [TODOの名前を表現する値オブジェクト]
	// バリデーションロジックは以下
	// - 名前は空ではないこと
	// - 名前は30文字以下であること

	value string
}

func NewName(value string) (error, Name) {
	empty := Name{}

	if value == "" {
		return errors.New("Name must not be empty"), empty
	} else if utf8.RuneCountInString(value) > 30 {
		return errors.New("Name too long"), empty
	} else {
		return nil, Name{value}
	}
}

func (name Name) Value() string {
	return name.value
}
