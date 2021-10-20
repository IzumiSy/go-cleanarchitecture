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
	id, _ := category.NewID(entity.GenerateID{})
	return Category{
		id:   id,
		name: name,
	}
}
