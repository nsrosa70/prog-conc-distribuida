package mutex

import (
	"concorrente/producerconsumer/event"
	"runtime"
	"sync"
)

type MutexEventBuffer struct {
	mu       sync.Mutex
	capacity int
	buffer   []event.Event
}

func NewMutexEventBuffer(capacity int) *MutexEventBuffer {
	return &MutexEventBuffer{
		mu:       sync.Mutex{},
		capacity: capacity,
		buffer:   []event.Event{},
	}
}

func (s *MutexEventBuffer) Add(e event.Event) {
	s.mu.Lock()
	for len(s.buffer) == s.capacity {
		s.mu.Unlock()
		runtime.Gosched()
		s.mu.Lock()
	}
	s.buffer = append(s.buffer, e)
	s.mu.Unlock()
}

func (s *MutexEventBuffer) Get() event.Event {
	s.mu.Lock()
	for len(s.buffer) == 0 {
		s.mu.Unlock()
		runtime.Gosched()
		s.mu.Lock()
	}
	ret := s.buffer[0]
	s.buffer = s.buffer[1:]
	s.mu.Unlock()

	return ret
}

