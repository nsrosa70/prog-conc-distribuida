package monitor

import (
	"aulas/concorrente/producerconsumer/event"
	"aulas/concorrente/producerconsumer/eventbuffer"
	"runtime"
	"sync"
)

type EventBufferMonitor interface {
	wait()
	signal()
	Add(event event.Event)
	Get() event.Event
}

type MonitorEventBuffer struct { // all variable are private, i.e., non-capital letter
	mu *sync.Mutex
	eb eventbuffer.EventBuffer
}

func NewMonitorEventBuffer(capacity, consumed int) *MonitorEventBuffer {
	eventBuffer := eventbuffer.EventBuffer{Capacity: capacity, Buffer: []event.Event{}, Consumed: consumed}
	return &MonitorEventBuffer{
		mu: &sync.Mutex{},
		eb: eventBuffer,
	}
}

func (b *MonitorEventBuffer) wait() {
	b.mu.Lock()
}
func (b *MonitorEventBuffer) signal() {
	b.mu.Unlock()
}

func (b *MonitorEventBuffer) Add(e event.Event) {
	b.wait()
	for len(b.eb.Buffer) == b.eb.Capacity {
		b.mu.Unlock()
		runtime.Gosched()
		b.mu.Lock()
	}
	b.eb.Buffer = append(b.eb.Buffer, e)
	b.signal()
}

func (b *MonitorEventBuffer) Get() event.Event {
	b.wait()
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
	b.signal()

	return e
}
