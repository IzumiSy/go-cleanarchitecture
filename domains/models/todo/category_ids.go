package todo

import (
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models/category"
)

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

func EmptyCategoryIds() CategoryIDs {
	return CategoryIDs{}
}

func (categoryIds CategoryIDs) Value() []category.ID {
	return categoryIds.value
}
