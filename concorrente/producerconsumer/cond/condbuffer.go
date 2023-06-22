package cond

import (
	"aulas/concorrente/producerconsumer/event"
	"aulas/concorrente/producerconsumer/eventbuffer"
	"sync"
)

type CondEventBuffer struct {
	cond *sync.Cond
	eb   eventbuffer.EventBuffer
}

func NewCondEventBuffer(capacity, consumed int) *CondEventBuffer {
	eventBuffer := eventbuffer.EventBuffer{Capacity: capacity, Buffer: []event.Event{}, Consumed: consumed}
	return &CondEventBuffer{
		cond: sync.NewCond(&sync.Mutex{}),
		eb:   eventBuffer,
	}
}

func (b *CondEventBuffer) Add(e event.Event) {
	b.cond.L.Lock()
	for len(b.eb.Buffer) == b.eb.Capacity {
		b.cond.Wait()
	}
	b.eb.Buffer = append(b.eb.Buffer, e)
	b.cond.Broadcast()
	b.cond.L.Unlock()
}

func (b *CondEventBuffer) Get() event.Event {
	b.cond.L.Lock()
	for len(b.eb.Buffer) == 0 {
		if b.eb.Consumed == 0 {
			b.cond.L.Unlock()
			return event.Event{E: ""}
		}
		b.cond.Wait()
	}
	e := b.eb.Buffer[0]
	b.eb.Buffer = b.eb.Buffer[1:]
	b.eb.Consumed--
	b.cond.Broadcast()
	b.cond.L.Unlock()

	return e
}
