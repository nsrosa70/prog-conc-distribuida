package channel

import "C"
import (
	"concorrente/producerconsumer/event"
)

type ChanEventBuffer struct {
	ch       chan bool
	capacity int
	buffer   []event.Event
}

func NewChanEventBuffer(capacity int) *ChanEventBuffer {
	return &ChanEventBuffer{
		ch:       make(chan bool, 1),
		capacity: capacity,
		buffer:   []event.Event{},
	}
}

func (eb *ChanEventBuffer) Add(e event.Event) {
	eb.ch <- true
	for len(eb.buffer) == eb.capacity {
		<- eb.ch
		//runtime.Gosched()
		eb.ch <- true
	}
	eb.buffer = append(eb.buffer, e)
	<- eb.ch
}

func (eb *ChanEventBuffer) Get() event.Event {
	eb.ch <- true
	for len(eb.buffer) == 0 {
		<- eb.ch
		//runtime.Gosched()
		eb.ch <- true
	}
	e := eb.buffer[0]
	eb.buffer = eb.buffer[1:]
	<- eb.ch

	return e
}

