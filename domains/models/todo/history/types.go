package history

import (
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models/entity"
)

// [HistoryエンティティのID]
type ID struct {
	entity.ID_
}

func NewID(builder entity.IDBuilder) (ID, errors.Domain) {
	id, err := builder.Build()
	return ID{ID_: id}, err
}

type CreatedAt struct {
	entity.Time_
}
