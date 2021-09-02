package todo

import (
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models/category"
	"go-cleanarchitecture/domains/models/entity"
	"unicode/utf8"
)

// [TodoエンティティのID]
type ID struct {
	entity.ID_
}

type CreatedAt struct {
	entity.Time_
}

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
		return empty, errors.Preconditional("Description must not be empty")
	} else if utf8.RuneCountInString(value) > 100 {
		return empty, errors.Preconditional("Description too long")
	} else {
		return Description{value}, errors.None
	}
}

func (description Description) Value() string {
	return description.value
}

type CategoryIDs struct {
	// [TODOにセットされるカテゴリの集合を表現する値オブジェクト]
	// バリデーションルールは以下
	// - セットできるカテゴリの上限は5個であること

	value []category.ID
}

func NewCategoryIds(categoryIds []category.ID) (CategoryIDs, errors.Domain) {
	empty := CategoryIDs{}

	if len(categoryIds) > 5 {
		return empty, errors.Preconditional("Too many categories")
	}

	return CategoryIDs{categoryIds}, errors.None
}

func EmptyCategoryIds() *CategoryIDs {
	return &CategoryIDs{}
}

func (categoryIds *CategoryIDs) Value() []category.ID {
	return categoryIds.value
}
