package semaforo

import (
	"runtime"
	"test/concorrente/producerconsumer/event"
	"test/concorrente/producerconsumer/eventbuffer"
)

type SemaforoEventBuffer struct {
	sem *Semaforo
	eb  eventbuffer.EventBuffer
}

func NewSemaforoEventBuffer(capacity, consumed int) *SemaforoEventBuffer {
	eventBuffer := eventbuffer.EventBuffer{Capacity: capacity, Buffer: []event.Event{}, Consumed: consumed}
	return &SemaforoEventBuffer{
		sem: NewSemaphore(1),
		eb:  eventBuffer,
	}
}

func (b *SemaforoEventBuffer) Add(e event.Event) {
	b.sem.P()
	for len(b.eb.Buffer) == b.eb.Capacity {
		b.sem.V()
		runtime.Gosched()
		b.sem.P()
	}
	b.eb.Buffer = append(b.eb.Buffer, e)
	b.sem.V()
}

func (b *SemaforoEventBuffer) Get() event.Event {
	b.sem.P()
	for len(b.eb.Buffer) == 0 {
		if b.eb.Consumed == 0 {
			b.sem.V()
			return event.Event{E: ""}
		}
		b.sem.V()
		runtime.Gosched()
		b.sem.P()
	}
	ret := b.eb.Buffer[0]
	b.eb.Buffer = b.eb.Buffer[1:]
	b.eb.Consumed--
	b.sem.V()

	return ret
}
