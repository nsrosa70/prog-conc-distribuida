package impl

import (
	"fmt"
	"sync"
)

type ISemaphore interface {
	P()
	V()
}

type Semaphore struct {
	n int32
	c *sync.Cond
}

func NewSemaphore(N int32) *Semaphore {
	return &Semaphore{
		n: N,
		c: sync.NewCond(new(sync.Mutex)),
	}
}

func (s *Semaphore) P() {
	s.c.L.Lock()
	for s.n <= 0 {
		s.c.Wait()
	}
	s.n--
	s.c.L.Unlock()
}

func (s *Semaphore) V() {
	s.c.L.Lock()
	s.n++
	fmt.Println(s.n)
	s.c.L.Unlock()
	s.c.Signal()
}
