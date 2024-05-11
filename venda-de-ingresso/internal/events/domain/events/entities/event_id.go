package entities

import "github.com/google/uuid"

type EventId struct {
	value uuid.UUID
}

func NewEventId(id string) (*EventId, error) {

	if id == "" {
		return &EventId{value: uuid.New()}, nil
	}
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, ErrInvalidId
	}

	return &EventId{value: uuid}, nil
}

func (props EventId) GetValue() uuid.UUID {
	return props.value
}
