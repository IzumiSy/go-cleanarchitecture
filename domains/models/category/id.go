package category

import (
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models/entity"
)

type Id entity.Id

func NewId(value string) (Id, errors.Domain) {
	_id, err := entity.NewId(value)
	return Id(_id), err
}

func (id Id) String() string {
	return id.String()
}

func GenerateId() Id {
	return Id(entity.GenerateId())
}
