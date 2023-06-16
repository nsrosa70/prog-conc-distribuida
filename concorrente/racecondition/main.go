package main

import (
	"fmt"
	"sync"
)

var count = 0

func increment(wg *sync.WaitGroup) {
	defer wg.Done()

	if count == 0 {
		count++
	}
}
func main() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go increment(&wg)

	wg.Add(1)
	go increment(&wg)

	wg.Wait()

	fmt.Println(count)
}
