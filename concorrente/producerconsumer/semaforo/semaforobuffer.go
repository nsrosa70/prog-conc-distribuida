package semaforo

import (
	"aulas/concorrente/producerconsumer/event"
	"runtime"
)

type SemaforoEventBuffer struct {
	sem      *Semaforo
	capacity int
	buffer   []event.Event
}

func NewSemaforoEventBuffer(capacity int) *SemaforoEventBuffer {
	return &SemaforoEventBuffer{
		sem:      NewSemaphore(1),
		capacity: capacity,
		buffer:   []event.Event{},
	}
}

func (s *SemaforoEventBuffer) Add(e event.Event) {
	s.sem.P()
	for len(s.buffer) == s.capacity {
		s.sem.V()
		runtime.Gosched()
		s.sem.P()
	}
	s.buffer = append(s.buffer, e)
	s.sem.V()
}

func (s *SemaforoEventBuffer) Get() event.Event {
	s.sem.P()
	for len(s.buffer) == 0 {
		s.sem.V()
		runtime.Gosched()
		s.sem.P()
	}
	ret := s.buffer[0]
	s.buffer = s.buffer[1:]
	s.sem.V()

	return ret
}
