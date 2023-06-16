package main

import (
	"fmt"
	"sync"
	"time"
)

func processors() {
	var x, y int
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		x = 1                   // A1
		fmt.Print("y:", y, " ") // A2
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		y = 1                   // B1
		fmt.Print("x:", x, " ") // B2
	}()

	wg.Wait()
}

func main(){
	for {
		processors()

		fmt.Println()
		time.Sleep(10 * time.Millisecond)
	}
}