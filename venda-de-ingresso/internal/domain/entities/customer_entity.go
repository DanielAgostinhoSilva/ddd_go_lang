package entities

import "encoding/json"

type CustomerConstructorProps struct {
	Id   string `json:"id"`
	Cpf  string `json:"cpf"`
	Name string `json:"name"`
}

type Customer struct {
	id   string
	cfp  string
	name string
}

func NewCustomer(cp *CustomerConstructorProps) *Customer {
	return &Customer{id: cp.Id, cfp: cp.Cpf, name: cp.Name}
}

func CreateCustomer(name, cpf string) *Customer {
	return NewCustomer(&CustomerConstructorProps{Name: name, Cpf: cpf})
}

func (c *Customer) ToJson() (string, error) {
	props := &CustomerConstructorProps{c.id, c.cfp, c.name}
	jsonData, err := json.Marshal(props)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func (c *Customer) Id() string {
	return c.id
}

func (c *Customer) Cfp() string {
	return c.cfp
}

func (c *Customer) Name() string {
	return c.name
}
