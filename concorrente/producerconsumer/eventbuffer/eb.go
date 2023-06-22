package eventbuffer

import (
	"aulas/concorrente/producerconsumer/event"
)

type EventBuffer struct {
	Primitive interface{}
	Capacity  int
	Buffer    []event.Event
	Consumed  int
}

type IEventBuffer interface {
	Get() event.Event
	Add(event2 event.Event)
}
