package models

import (
	"go-cleanarchitecture/domains/models/category"
)

type Category struct {
	// [TODOに紐づくカテゴリを表現するエンティティ

	id   category.ID
	name category.Name
}

func NewCategory(name category.Name) Category {
	return Category{
		id:   category.GenerateID(),
		name: name,
	}
}
