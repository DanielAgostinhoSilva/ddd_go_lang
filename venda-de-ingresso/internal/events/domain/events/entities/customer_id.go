package entities

import (
	"errors"
	"github.com/google/uuid"
)

var (
	ErrInvalidId = errors.New("invalid value")
)

type CustomerId struct {
	value uuid.UUID
}

func (c *CustomerId) Value() uuid.UUID {
	return c.value
}

func NewCustomerId(id string) (*CustomerId, error) {
	if id == "" {
		return &CustomerId{value: uuid.New()}, nil
	}
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, ErrInvalidId
	}

	return &CustomerId{value: uuid}, nil
}
