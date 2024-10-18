package queue

import (
	"runtime"
	"sync"
	"test/mymom/services/messagingservice/event"
)

type MutexQueue struct {
	mu       sync.Mutex
	capacity int
	queue    []event.Event
}

func (s *MutexQueue) Push(i event.Event) {

	s.mu.Lock()
	for len(s.queue) == s.capacity {
		s.mu.Unlock()
		runtime.Gosched()
		s.mu.Lock()
	}
	s.queue = append(s.queue, i)
	s.mu.Unlock()
}

func (s *MutexQueue) Pop() event.Event {

	s.mu.Lock()
	for len(s.queue) == 0 {
		s.mu.Unlock()
		runtime.Gosched()
		s.mu.Lock()
	}
	r := s.queue[0]
	s.queue = s.queue[1:]
	s.mu.Unlock()

	return r
}

func (s MutexQueue) Size() int {
	return len(s.queue)
}

func NewMutexQueue(capacity int) *MutexQueue {
	return &MutexQueue{
		mu:       sync.Mutex{},
		capacity: capacity,
		queue:    []event.Event{},
	}
}
