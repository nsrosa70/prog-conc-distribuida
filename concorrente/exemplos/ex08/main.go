/* Exemplo gerado pelo ChatGPT */
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var cond = sync.NewCond(&sync.Mutex{})
	var ready bool

	wg.Add(1)
	go func() {
		defer wg.Done()

		fmt.Println("Goroutine waiting...")
		cond.L.Lock()
		for !ready {
			cond.Wait()
		}
		fmt.Println("Goroutine woke up!")
		cond.L.Unlock()
	}()
	// Simulating some work
	fmt.Println("Main thread working for 5 seconds...")
	time.Sleep(5 * time.Second)
	cond.L.Lock()
	ready = true
	cond.Signal()
	cond.L.Unlock()
	fmt.Println("Main thread finished!")
	wg.Wait()
}
