package mutexqueue

import (
	"runtime"
	"sync"
)

type MutexQueue struct {
	mu       sync.Mutex
	capacity int
	queue    []int
}

func NewMutexQueue(capacity int) *MutexQueue {
	return &MutexQueue{
		mu:       sync.Mutex{},
		capacity: capacity,
		queue:    []int{},
	}
}

func (s *MutexQueue) Push(i int, wg *sync.WaitGroup) {
	defer wg.Done()

	s.mu.Lock()
	for len(s.queue) == s.capacity {
		s.mu.Unlock()
		runtime.Gosched()
		s.mu.Lock()
	}
	s.queue = append(s.queue, i)
	s.mu.Unlock()
}

func (s *MutexQueue) Pop(wg *sync.WaitGroup) int {
	defer wg.Done()

	s.mu.Lock()
	for len(s.queue) == 0 {
		s.mu.Unlock()
		runtime.Gosched()
		s.mu.Lock()
	}
	ret := s.queue[0]
	s.queue = s.queue[1:]
	s.mu.Unlock()

	return ret
}

