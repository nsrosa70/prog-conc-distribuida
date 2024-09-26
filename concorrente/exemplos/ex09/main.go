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

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() { // G1
			defer wg.Done()

			fmt.Println("Goroutine 1 waiting...")
			cond.L.Lock()
			for !ready {
				cond.Wait()
			}
			fmt.Println("Goroutine 1 woke up!")
			cond.L.Unlock()
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()

			// Simulating some work
			fmt.Println("Goroutine 2 working for 5 seconds...")
			time.Sleep(5 * time.Second)
			cond.L.Lock()
			ready = true
			cond.Signal()
			cond.L.Unlock()
			fmt.Println("Go routine 2 finished!")
		}()
		wg.Wait()
	}
}
