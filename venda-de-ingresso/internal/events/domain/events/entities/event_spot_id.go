package entities

import "github.com/google/uuid"

type EventSpotId struct {
	value uuid.UUID
}

func NewEventSpotId(id string) (*EventSpotId, error) {
	if id == "" {
		return &EventSpotId{value: uuid.New()}, nil
	}
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, ErrInvalidId
	}

	return &EventSpotId{value: uuid}, nil
}

func (props EventSpotId) GetValue() uuid.UUID {
	return props.value
}
