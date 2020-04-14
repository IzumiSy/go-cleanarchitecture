package todo

import (
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models/category"
)

type CategoryIds struct {
	// [TODOにセットされるカテゴリの集合を表現する値オブジェクト]
	// バリデーションルールは以下
	// - セットできるカテゴリの上限は5個であること

	value []category.Id
}

func NewCategoryIds(categoryIds []category.Id) (CategoryIds, errors.Domain) {
	empty := CategoryIds{}

	if len(categoryIds) > 5 {
		return empty, errors.Invalid("Too many categories")
	}

	return CategoryIds{categoryIds}, errors.None
}

func EmptyCategoryIds() CategoryIds {
	return CategoryIds{}
}
