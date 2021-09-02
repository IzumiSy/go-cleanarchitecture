package todo

import (
	"go-cleanarchitecture/domains/models/entity"
	"go-cleanarchitecture/domains/models/todo/history"
)

type HistoryType struct {
	_type string
}

func (t HistoryType) String() string {
	return t._type
}

var (
	TodoCreated        HistoryType = HistoryType{_type: "todo_created"}
	NameUpdated        HistoryType = HistoryType{_type: "name_updated"}
	DescriptionUpdated HistoryType = HistoryType{_type: "description_updated"}
)

type valueable interface {
	Value() string
}

type History struct {
	// [TODOの変更履歴を表現するエンティティ]

	id          history.ID
	todoID      ID
	historyType HistoryType
	previous    valueable
	current     valueable
	createdAt   history.CreatedAt
}

func (h History) ID() history.ID {
	return h.id
}

func (h History) Type() HistoryType {
	return h.historyType
}

func (h History) Previous() valueable {
	return h.previous
}

func (h History) Current() valueable {
	return h.current
}

func (h History) CreatedAt() history.CreatedAt {
	return h.createdAt
}

type Histories struct {
	// [あるTODOの変更履歴を集合として表現するエンティティ]

	todoID            ID
	value             []History
	lastUpdatesByType map[HistoryType]valueable
}

func NewHistories(todoID ID) *Histories {
	histories := &Histories{
		todoID:            todoID,
		lastUpdatesByType: map[HistoryType]valueable{},
	}
	histories.AddHistory(TodoCreated, nil)
	return histories
}

// 新しい変更履歴を追加する
func (h *Histories) AddHistory(historyType HistoryType, value valueable) {
	previous := h.lastUpdatesByType[historyType]

	h.value = append(h.value, History{
		id:          history.ID{ID_: entity.GenerateID()},
		todoID:      h.todoID,
		historyType: historyType,
		previous:    previous,
		current:     value,
		createdAt:   history.CreatedAt{Time_: entity.GenerateTime()},
	})

	h.lastUpdatesByType[historyType] = value
}

func (h *Histories) Last() *History {
	if len(h.value) == 0 {
		return nil
	}
	return &h.value[len(h.value)-1]
}

func (h *Histories) At(index int) *History {
	if len(h.value) < index {
		return nil
	}
	return &h.value[index]
}

func (h *Histories) Length() int {
	return len(h.value)
}
