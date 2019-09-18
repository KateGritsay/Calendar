
//go:generate protoc --go_out=. calendarpb.proto


package calendar

import (
	"errors"
	"fmt"
	calendar "github.com/KateGritsay/Calendar/pkg/calendarpb"
)


type Calendar struct {
	events map[uint64]calendar.Event
}

func NewCalendar() Calendar {
	cal := Calendar{events: make(map[uint64]calendar.Event)}
	return cal
}

func (c *Calendar) AddEvent(event calendar.Event) calendar.Event {
	var counter uint64
	counter++
	event.UUID = counter
	c.events[event.UUID] = event

	return event
}

func (c *Calendar) DeleteEvent(event calendar.Event) error {
	_, ok := c.events[event.UUID]
	if ok {
		delete(c.events, event.UUID)
	}
	return errors.New(fmt.Sprintf("Can't delete event: UUID %d not found", event.UUID))

}

func (c *Calendar) ReplaceEvent(event calendar.Event) error {
	_, ok := c.events[event.UUID]
	if ok {
		c.events[event.UUID] = event
	}
	return errors.New(fmt.Sprintf("UUID %d not found", event.UUID))
}
