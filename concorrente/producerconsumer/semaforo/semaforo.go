package semaforo

import "sync"

type ISemaforo interface {
	P()
	V()
}

type Semaforo struct {
	n int32
	c *sync.Cond
}

func NewSemaphore(N int32) *Semaforo {
	return &Semaforo{
		n: N,
		c: sync.NewCond(new(sync.Mutex)),
	}
}

func (s *Semaforo) P() {
	s.c.L.Lock()
	for s.n <= 0 {
		s.c.Wait()
	}
	s.n--
	s.c.L.Unlock()
}

func (s *Semaforo) V() {
	s.c.L.Lock()
	s.n++
	s.c.L.Unlock()
	s.c.Signal()
}


