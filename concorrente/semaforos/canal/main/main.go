package main

import (
	"aulas/concorrente/semaforos/canal/impl"
	"fmt"
	"sync"
	"time"
)

var i int
var s impl.Semaphore
var ws sync.WaitGroup

func increment() {
	defer ws.Done()
	s.P()
	i = i + 1
	s.V()
}

func decrement() {
	defer ws.Done()
	s.P()
	i = i - 1
	s.V()
}

func main() {
	s = impl.NewSemaphore(1)

	t1 := time.Now()
	for i := 0; i < 100000; i++ {
		ws.Add(1)
		go increment()

		ws.Add(1)
		go decrement()
	}
	fmt.Println(time.Now().Sub(t1))

	ws.Wait()

	fmt.Println(i)
}
