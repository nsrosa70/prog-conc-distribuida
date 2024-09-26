package impl

import "sync"

type ISemaphore interface {
	P()
	V()
}

type Semaphore struct {
	n int32
	m sync.Mutex
}

func NewSemaphore(N int32) Semaphore {
	s := Semaphore{n: N, m: *new(sync.Mutex)}
	return s
}

func (s *Semaphore) P() { // Decrement / Wait
	s.m.Lock()
}

func (s *Semaphore) V() { // Increment Signal
	s.m.Unlock()
}
