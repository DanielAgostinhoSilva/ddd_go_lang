package value_objects

import (
	"errors"
	"reflect"
)

var (
	ErrNameIsRequired = errors.New("name is required")
)

type Name struct {
	name string
}

func (n *Name) Name() string {
	return n.name
}

func NewName(name string) (*Name, error) {
	newName := &Name{name: name}
	if !newName.IsValid() {
		return nil, ErrNameIsRequired
	}
	return newName, nil
}

func (n *Name) Equals(name *Name) bool {
	return reflect.DeepEqual(n, name)
}

func (n *Name) IsValid() bool {
	return len(n.name) >= 3
}
