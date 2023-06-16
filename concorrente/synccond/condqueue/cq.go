package condqueue

import "sync"

type CondQueue struct {
	cond     *sync.Cond
	capacity int
	queue    []int
}

func NewCondQueue(capacity int) *CondQueue {
	return &CondQueue{
		cond:     sync.NewCond(&sync.Mutex{}),
		capacity: capacity,
		queue:    []int{},
	}
}

func (s *CondQueue) Push(i int, wg *sync.WaitGroup) {
	wg.Done()
	// Acquire lock before entering the critical section
	s.cond.L.Lock()
	for len(s.queue) == s.capacity {
		// Wait for a signal sent by Broadcast()
		// When receives a signal, it goes to the head of the loop
		// then checks the condition again
		s.cond.Wait()
	}

	s.queue = append(s.queue, i)

	// Because condition (= length of s.queue) is changed,
	// it sends a signal to all the goroutines
	// Because they wait for the signal, it doesn't enter busy-loop,
	// so it is more efficient.
	s.cond.Broadcast()
	s.cond.L.Unlock()
}

func (s *CondQueue) Pop(wg *sync.WaitGroup) int {
	defer wg.Done()

	s.cond.L.Lock()
	for len(s.queue) == 0 {
		s.cond.Wait()
	}

	ret := s.queue[0]
	s.queue = s.queue[1:]
	s.cond.Broadcast()
	s.cond.L.Unlock()

	return ret
}

