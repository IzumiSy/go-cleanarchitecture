package models

import (
	"go-cleanarchitecture/domains/models/category"
	"go-cleanarchitecture/domains/models/entity"
)

type Category struct {
	// [TODOに紐づくカテゴリを表現するエンティティ

	id   category.ID
	name category.Name
}

func NewCategory(name category.Name) Category {
	return Category{
		id:   category.ID{ID_: entity.GenerateID()},
		name: name,
	}
}
