package channel

import (
	"runtime"
	"test/concorrente/producerconsumer/event"
	"test/concorrente/producerconsumer/eventbuffer"
)

type ChanEventBuffer struct {
	ch chan bool
	eb eventbuffer.EventBuffer
}

func NewChanEventBuffer(capacity int, consumed int) *ChanEventBuffer {
	eventBuffer := eventbuffer.EventBuffer{Capacity: capacity, Buffer: []event.Event{}, Consumed: consumed}
	return &ChanEventBuffer{
		ch: make(chan bool, 1),
		eb: eventBuffer,
	}
}

func (b *ChanEventBuffer) Add(e event.Event) {
	b.ch <- true
	for len(b.eb.Buffer) == b.eb.Capacity {
		<-b.ch
		runtime.Gosched()
		b.ch <- true
	}
	b.eb.Buffer = append(b.eb.Buffer, e)
	<-b.ch
}

func (b *ChanEventBuffer) Get() event.Event {
	b.ch <- true
	for len(b.eb.Buffer) == 0 {
		if b.eb.Consumed == 0 {
			<-b.ch
			return event.Event{E: ""}
		}
		<-b.ch
		runtime.Gosched()
		b.ch <- true
	}
	e := b.eb.Buffer[0]
	b.eb.Buffer = b.eb.Buffer[1:]
	b.eb.Consumed--
	<-b.ch

	return e
}
