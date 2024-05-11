package entities

import "github.com/google/uuid"

type EventSectionId struct {
	value uuid.UUID
}

func NewEventSectionId(id string) (*EventSectionId, error) {
	if id == "" {
		return &EventSectionId{value: uuid.New()}, nil
	}
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, ErrInvalidId
	}

	return &EventSectionId{value: uuid}, nil
}

func (props EventSectionId) GetValue() uuid.UUID {
	return props.value
}
