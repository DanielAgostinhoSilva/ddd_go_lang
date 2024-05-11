package entities

import "encoding/json"

type PartnerProps struct {
	id   string
	name string
}

type Partner struct {
	id   *PartnerId
	name string
}

func NewPartner(props *PartnerProps) (*Partner, error) {
	id, err := NewPartnerId(props.id)
	if err != nil {
		return nil, err
	}
	return &Partner{id: id, name: props.name}, nil
}

func CreatePartner(command *CreatePartnerCommand) (*Partner, error) {
	return NewPartner(&PartnerProps{name: command.Name})
}

func (p *Partner) ToJson() (string, error) {
	props := &PartnerProps{p.id.value.String(), p.name}
	jsonData, err := json.Marshal(props)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func (p *Partner) Equals(partner *Partner) bool {
	return p.id.Value() == partner.id.Value()
}

func (p *Partner) Id() *PartnerId {
	return p.id
}

func (p *Partner) Name() string {
	return p.name
}
