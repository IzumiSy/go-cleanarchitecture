package user

import (
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models/entity"
)

// [UserのIDを表現する値オブジェクト]

type Id entity.Id

func New(value string) (Id, errors.Domain) {
	_id, err := entity.NewId(value)
	return Id(_id), err
}

func (id Id) String() string {
	return id.String()
}

func GenerateId() Id {
	return Id(entity.GenerateId())
}
