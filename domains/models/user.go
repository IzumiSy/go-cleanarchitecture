package models

import (
	"go-cleanarchitecture/domains/models/entity"
	"go-cleanarchitecture/domains/models/user"
)

type User struct {
	// Userを表現するエンティティ

	id   user.ID
	name user.Name
}

func NewUser(name user.Name) User {
	return User{
		id:   user.ID{ID_: entity.GenerateID()},
		name: name,
	}
}

func BuildUser(id user.ID, name user.Name) User {
	return User{
		id:   id,
		name: name,
	}
}

func (user User) ID() user.ID {
	return user.id
}

func (user User) Name() user.Name {
	return user.name
}
