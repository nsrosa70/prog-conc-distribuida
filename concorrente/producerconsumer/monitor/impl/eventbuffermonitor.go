package monitor

import (
	"aulas/concorrente/producerconsumer/event"
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
	mu       *sync.Mutex
	capacity int
	buffer   []event.Event
}

func NewMonitorEventBuffer(capacity int) *MonitorEventBuffer {
	return &MonitorEventBuffer{
		mu:       &sync.Mutex{},
		capacity: capacity,
		buffer:   []event.Event{},
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
	for len(b.buffer) == b.capacity {
		b.mu.Unlock()
		runtime.Gosched()
		b.mu.Lock()
	}
	b.buffer = append(b.buffer, e)
	b.signal()
}

func (b *MonitorEventBuffer) Get() event.Event {
	b.wait()
	for len(b.buffer) == 0 {
		b.mu.Unlock()
		runtime.Gosched()
		b.mu.Lock()
	}
	e := b.buffer[0]
	b.buffer = b.buffer[1:]
	b.signal()

	return e
}
