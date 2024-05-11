package entities

import (
	"encoding/json"
	value_objects "github.com/DanielAgostinhoSilva/ddd_go_lang/venda-de-ingressos/internal/common/domain/value-objects"
)

type CustomerConstructorProps struct {
	Id   string `json:"value"`
	Cpf  string `json:"cpf"`
	Name string `json:"name"`
}

type Customer struct {
	id   *CustomerId
	cpf  *value_objects.Cpf
	name *value_objects.Name
}

func NewCustomer(cp *CustomerConstructorProps) (*Customer, error) {
	cpf, err := value_objects.NewCpf(cp.Cpf)
	if err != nil {
		return nil, err
	}
	name, err := value_objects.NewName(cp.Name)
	if err != nil {
		return nil, err
	}
	newId, err := NewCustomerId(cp.Id)
	if err != nil {
		return nil, err
	}

	return &Customer{id: newId, cpf: cpf, name: name}, nil
}

func CreateCustomer(name, cpf string) (*Customer, error) {
	return NewCustomer(&CustomerConstructorProps{Name: name, Cpf: cpf})
}

func (c *Customer) ToJson() (string, error) {
	props := &CustomerConstructorProps{c.id.value.String(), c.cpf.Value(), c.name.Value()}
	jsonData, err := json.Marshal(props)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func (c *Customer) Equals(id *CustomerId) bool {
	return c.id == id
}

func (c *Customer) Id() *CustomerId {
	return c.id
}

func (c *Customer) Cfp() *value_objects.Cpf {
	return c.cpf
}

func (c *Customer) Name() *value_objects.Name {
	return c.name
}
