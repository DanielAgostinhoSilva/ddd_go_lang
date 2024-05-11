package entities

import (
	value_objects "github.com/DanielAgostinhoSilva/ddd_go_lang/venda-de-ingressos/internal/common/domain/value-objects"
	"github.com/google/uuid"
	"time"
)

type CreateEventCommand struct {
	Name        string
	Description string
	Date        time.Time
	PartnerId   string
}

type AddSectionCommand struct {
	Name        string
	Description string
	TotalSpot   int
	Price       float64
}

type EventProps struct {
	Id                string
	Name              string
	Description       string
	Date              time.Time
	Published         bool
	TotalSpots        int
	TotalSpotReserved int
	PartnerId         string
	Sections          []EventSection
}

type Event struct {
	id                *EventId
	name              value_objects.Name
	description       string
	date              time.Time
	published         bool
	totalSpots        int
	totalSpotReserved int
	partnerId         *PartnerId
	sections          []EventSection
}

func NewEvent(props *EventProps) (*Event, error) {
	name, err := value_objects.NewName(props.Name)
	if err != nil {
		return nil, err
	}
	partnerId, err := NewPartnerId(props.PartnerId)
	if err != nil {
		return nil, err
	}
	eventId, err := NewEventId(props.Id)
	if err != nil {
		return nil, err
	}

	return &Event{
		id:                eventId,
		name:              *name,
		description:       props.Description,
		date:              props.Date,
		published:         props.Published,
		totalSpots:        props.TotalSpots,
		totalSpotReserved: props.TotalSpotReserved,
		partnerId:         partnerId,
		sections:          props.Sections,
	}, err
}

func CreateEvent(command CreateEventCommand) (*Event, error) {
	props := &EventProps{
		Name:        command.Name,
		Description: command.Description,
		Date:        command.Date,
		PartnerId:   command.PartnerId,
	}
	event, err := NewEvent(props)
	return event, err
}

func (props *Event) Publish() {
	props.published = true
}

func (props *Event) UnPublish() {
	props.published = false
}

func (props *Event) AddSection(command AddSectionCommand) error {
	createEventSectionCommand := &CreateEventSectionCommand{
		command.Name,
		command.Description,
		command.TotalSpot,
		command.Price,
	}
	eventSection, err := CreateEventSection(createEventSectionCommand)
	if eventSection != nil {
		props.sections = append(props.sections, *eventSection)
		props.totalSpots += eventSection.totalSpot
	}
	return err
}

func (props *Event) ChangeName(name string) error {
	newName, err := value_objects.NewName(name)
	if newName != nil {
		props.name = *newName
	}
	return err
}

func (props *Event) GetId() uuid.UUID {
	return props.id.GetValue()
}

func (props *Event) GetName() string {
	return props.name.Value()
}

func (props *Event) GetDescription() string {
	return props.description
}

func (props *Event) GetDate() time.Time {
	return props.date
}

func (props *Event) IsPublished() bool {
	return props.published
}

func (props *Event) GetTotalSpots() int {
	return props.totalSpots
}

func (props *Event) GetTotalSpotReserved() int {
	return props.totalSpotReserved
}

func (props *Event) GetPartnerId() *PartnerId {
	return props.partnerId
}

func (props *Event) GetSections() []EventSection {
	return props.sections
}
