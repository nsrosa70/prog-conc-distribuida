package impl

type ISemaphore interface {
	P()
	V()
}

type Semaphore chan struct{}

func NewSemaphore(N int32) Semaphore {
	s := Semaphore(make(chan struct{}, N))
	for i := N; i > 0; i-- {
		s.V()
	}
	return s
}

func (s Semaphore) P() { // Decrement / Wait
	<-s
}

func (s Semaphore) V() { // Increment Signal
	s <- struct{}{}
}
