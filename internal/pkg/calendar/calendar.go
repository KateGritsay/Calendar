
//go:generate protoc --go_out=. calendar.proto


package calendar

import (
	"errors"
	"fmt"
	""
)

type Calendar struct {
	events map[uint64]Event
}

func NewCalendar() Calendar {
	cal := Calendar{events: make(map[uint64]Event)}
	return cal
}

func (cal *Calendar) AddEvent(event Event) Event {
	var counter uint64
	counter++
	event.UUID = counter
	cal.events[event.UUID] = event

	return event
}

func (cal *Calendar) DeleteEvent(event Event) error {
	_, ok := cal.events[event.UUID]
	if ok {
		delete(cal.events, event.UUID)
	}
	return errors.New(fmt.Sprintf("Can't delete event: UUID %d not found", event.UUID))

}

func (cal *Calendar) ReplaceEvent(event Event) error {
	_, ok := cal.events[event.UUID]
	if ok {
		cal.events[event.UUID] = event
	}
	return errors.New(fmt.Sprintf("UUID %d not found", event.UUID))
}
