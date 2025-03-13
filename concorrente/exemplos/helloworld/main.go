// Created by chatgpt

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// Create a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Define the number of goroutines to create
	numGoroutines := 5

	for i := 1; i <= numGoroutines; i++ {
		// Increment the WaitGroup to indicate a new goroutine
		wg.Add(1)

		// Goroutine for printing numbers
		go func(i int) {
			defer wg.Done() // Decrement the WaitGroup when the goroutine finishes
			fmt.Printf("Goroutine %d: Hello, World!\n", i)
		}(i)
	}

	runtime.GOMAXPROCS(1)

	fmt.Println(">>>", runtime.NumCPU())

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("All goroutines have finished.")
}
