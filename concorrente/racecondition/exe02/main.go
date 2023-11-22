package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	j := 0
	for i := 0; i < 1; i++ {
		wg.Add(1)
		defer wg.Done()
		go func() {
			fmt.Println(j)
			j++
		}()
	}
	wg.Wait() //original
}
