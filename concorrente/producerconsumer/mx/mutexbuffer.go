package mx

import (
	"aulas/concorrente/producerconsumer/event"
	"aulas/concorrente/producerconsumer/eventbuffer"
	"runtime"
	"sync"
)

type MutexEventBuffer struct {
	mu sync.Mutex
	eb eventbuffer.EventBuffer
}

func NewMutexEventBuffer(capacity, consumed int) *MutexEventBuffer {
	eventBuffer := eventbuffer.EventBuffer{Capacity: capacity, Buffer: []event.Event{}, Consumed: consumed}
	return &MutexEventBuffer{
		mu: sync.Mutex{},
		eb: eventBuffer,
	}
}

func (b *MutexEventBuffer) Add(e event.Event) {
	b.mu.Lock()
	for len(b.eb.Buffer) == b.eb.Capacity {
		b.mu.Unlock()
		runtime.Gosched()
		b.mu.Lock()
	}
	b.eb.Buffer = append(b.eb.Buffer, e)
	b.mu.Unlock()
}

func (b *MutexEventBuffer) Get() event.Event {
	b.mu.Lock()
	for len(b.eb.Buffer) == 0 {
		if b.eb.Consumed == 0 {
			b.mu.Unlock()
			return event.Event{E: ""}
		}
		b.mu.Unlock()
		runtime.Gosched()
		b.mu.Lock()
	}
	e := b.eb.Buffer[0]
	b.eb.Buffer = b.eb.Buffer[1:]
	b.eb.Consumed--
	b.mu.Unlock()

	return e
}
