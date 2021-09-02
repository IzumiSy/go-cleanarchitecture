package history

import (
	"go-cleanarchitecture/domains/models/entity"
)

// [HistoryエンティティのID]
type ID struct {
	entity.ID_
}

type CreatedAt struct {
	entity.Time_
}
