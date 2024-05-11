package value_objects

import (
	"errors"
	"github.com/google/uuid"
	"reflect"
)

var (
	ErrInvalidId = errors.New("invalid id")
)

type ID struct {
	value uuid.UUID
}

func NewID(id string) (*ID, error) {
	if id == "" {
		return &ID{value: uuid.New()}, nil
	}
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, ErrInvalidId
	}

	return &ID{value: uuid}, nil
}

func (i *ID) Equals(id *ID) bool {
	return reflect.DeepEqual(i, id)
}

func (i *ID) Value() uuid.UUID {
	return i.value
}
