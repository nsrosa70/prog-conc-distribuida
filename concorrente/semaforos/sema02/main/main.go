package main

import (
	"fmt"
	"semaphores/semaphore02/impl"
	"sync"
)

var i int
var s impl.Semaphore
var ws sync.WaitGroup

func increment(){
	defer ws.Done()
	s.P()
    i = i + 1
    s.V()
}

func decrement(){
	defer ws.Done()
	s.P()
	i = i - 1
	s.V()
}

func main(){
	s = impl.NewSemaphore(1)

	for i := 0; i < 100; i ++ {
		ws.Add(1)
		go increment()

		ws.Add(1)
		go decrement()
	}

	ws.Wait()

	fmt.Println(i)
}
