package channel

import (
	"aulas/concorrente/producerconsumer/event"
	"runtime"
	"sync/atomic"
)

type ChanEventBuffer struct {
	ch       chan bool
	capacity int
	buffer   []event.Event
	size     int32
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
		<-eb.ch
		runtime.Gosched()
		eb.ch <- true
	}
	eb.buffer = append(eb.buffer, e)
	eb.size++
	<-eb.ch
}

func (eb *ChanEventBuffer) Get() event.Event {
	eb.ch <- true
	for len(eb.buffer) == 0 {
		<-eb.ch
		runtime.Gosched()
		eb.ch <- true
	}
	e := eb.buffer[0]
	eb.buffer = eb.buffer[1:]
	eb.size--
	<-eb.ch

	return e
}

func (eb *ChanEventBuffer) Size() int32 {
	return atomic.LoadInt32(&eb.size)
}
