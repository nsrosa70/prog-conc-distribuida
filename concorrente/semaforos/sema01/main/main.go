package main

import (
	"concorrente/semaforos/sema01/impl"
	"fmt"
	"sync"
)

var n int

var s impl.Semaphore
var ws sync.WaitGroup

func increment(){
	defer ws.Done()

	s.P()
    n = n + 1
    s.V()
}

func decrement(){
	defer ws.Done()

	s.P()
	n = n - 1
	s.V()
}

func main(){
	s = *impl.NewSemaphore(1)

	for i := 0; i < 100; i ++ {
		ws.Add(1)
		go increment()

		ws.Add(1)
		go decrement()
	}

	ws.Wait()

	fmt.Println(n)
}

