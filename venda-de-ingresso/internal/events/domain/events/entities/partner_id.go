package entities

import (
	"errors"
	"github.com/google/uuid"
)

var (
	ErrInvalidPartnerId = errors.New("invalid partnerId")
)

type PartnerId struct {
	value uuid.UUID
}

func (c *PartnerId) Value() uuid.UUID {
	return c.value
}

func NewPartnerId(id string) (*PartnerId, error) {
	if id == "" {
		return &PartnerId{value: uuid.New()}, nil
	}
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, ErrInvalidPartnerId
	}

	return &PartnerId{value: uuid}, nil
}
