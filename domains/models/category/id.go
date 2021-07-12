package category

import (
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models/entity"
)

type ID struct {
	// カテゴリのIDを表現する値オブジェクト

	value entity.ID
}

func NewID(value string) (ID, errors.Domain) {
	id, err := entity.NewID(value)
	return ID{id}, err
}

func (id ID) Value() string {
	return id.value.String()
}

func GenerateID() ID {
	return ID{entity.GenerateID()}
}
