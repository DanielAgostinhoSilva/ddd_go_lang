package value_objects

import (
	"errors"
	"reflect"
)

var (
	ErrNameIsRequired = errors.New("value is required")
)

type Name struct {
	value string
}

func (n *Name) Value() string {
	return n.value
}

func NewName(name string) (*Name, error) {
	newName := &Name{value: name}
	if !newName.IsValid() {
		return nil, ErrNameIsRequired
	}
	return newName, nil
}

func (n *Name) Equals(name *Name) bool {
	return reflect.DeepEqual(n, name)
}

func (n *Name) IsValid() bool {
	return len(n.value) >= 3
}
