package cond

import "C"
import (
	"multithread/producerconsumer/event"
	"sync"
)

type CondEventBuffer struct {
	cond     *sync.Cond
	capacity int
	buffer    []event.Event
}

func NewCondEventBuffer(capacity int) *CondEventBuffer {
	return &CondEventBuffer{
		cond:     sync.NewCond(&sync.Mutex{}),
		capacity: capacity,
		buffer:    []event.Event{},
	}
}

func (eb *CondEventBuffer) Add(e event.Event) {
	eb.cond.L.Lock()
	for len(eb.buffer) == eb.capacity {
		eb.cond.Wait()
	}
	eb.buffer = append(eb.buffer, e)
	eb.cond.Broadcast()
	eb.cond.L.Unlock()
}

func (eb *CondEventBuffer) Get() event.Event {
	eb.cond.L.Lock()
	for len(eb.buffer) == 0 {
		eb.cond.Wait()
	}
	e := eb.buffer[0]
	eb.buffer = eb.buffer[1:]
	eb.cond.Broadcast()
	eb.cond.L.Unlock()

	return e
}

