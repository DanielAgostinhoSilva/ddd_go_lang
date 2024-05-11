package entities

import (
	"errors"
	value_objects "github.com/DanielAgostinhoSilva/ddd_go_lang/venda-de-ingressos/internal/common/domain/value-objects"
	"github.com/google/uuid"
)

var (
	ErrInvalidPrice = errors.New("Invalid Price")
)

type EventSectionProps struct {
	Id                string
	Name              string
	Description       string
	Published         bool
	TotalSpot         int
	TotalSpotReserved int
	Price             float64
	Spots             []EventSpot
}

type EventSection struct {
	id                *EventSectionId
	name              *value_objects.Name
	description       string
	published         bool
	totalSpot         int
	totalSpotReserved int
	price             *value_objects.Price
	spots             []EventSpot
}

type CreateEventSectionCommand struct {
	Name        string
	Description string
	TotalSpot   int
	Price       float64
}

func NewEventSection(props *EventSectionProps) (*EventSection, error) {
	id, err := NewEventSectionId(props.Id)
	if err != nil {
		return nil, err
	}
	name, err := value_objects.NewName(props.Name)
	if err != nil {
		return nil, err
	}
	price, err := value_objects.NewPrice(props.Price)
	if err != nil {
		return nil, err
	}

	return &EventSection{
		id,
		name,
		props.Description,
		props.Published,
		props.TotalSpot,
		props.TotalSpotReserved,
		price,
		props.Spots,
	}, nil
}

func CreateEventSection(command *CreateEventSectionCommand) (*EventSection, error) {
	eventSection, err := NewEventSection(&EventSectionProps{
		Id:                "",
		Name:              command.Name,
		Description:       command.Description,
		TotalSpot:         command.TotalSpot,
		TotalSpotReserved: 0,
		Price:             command.Price,
		Spots:             []EventSpot{},
	})

	if err != nil {
		return nil, err
	}

	err = eventSection.initSpots()

	return eventSection, err
}

func (props *EventSection) initSpots() error {
	for i := 0; i < props.totalSpot; i++ {
		eventSpot, err := CreateEventSpot()
		if err != nil {
			return err
		}
		props.spots = append(props.spots, *eventSpot)
	}
	return nil
}

func (props *EventSection) ChangeName(name string) error {
	newName, err := value_objects.NewName(name)
	if newName != nil {
		props.name = newName
	}
	return err
}

func (props *EventSection) ChangeDescription(description string) {
	props.description = description
}

func (props *EventSection) ChangePrice(price float64) error {
	newPrice, err := value_objects.NewPrice(price)
	if newPrice != nil {
		props.price = newPrice
	}
	return err
}

func (props *EventSection) Publish() {
	props.published = true
}

func (props *EventSection) UnPublish() {
	props.published = false
}

func (props *EventSection) PublishAll() {
	props.Publish()
	for index := range props.spots {
		props.spots[index].Publish()
	}
}

func (props *EventSection) GetId() uuid.UUID {
	return props.id.GetValue()
}

func (props *EventSection) GetName() string {
	return props.name.Value()
}

func (props *EventSection) GetDescription() string {
	return props.description
}

func (props *EventSection) IsPublished() bool {
	return props.published
}

func (props *EventSection) GetTotalSpot() int {
	return props.totalSpot
}

func (props *EventSection) GetTotalSpotReserved() int {
	return props.totalSpotReserved
}

func (props *EventSection) GetPrice() float64 {
	return props.price.GetValue()
}

func (props *EventSection) GetSpots() []EventSpot {
	return props.spots
}
