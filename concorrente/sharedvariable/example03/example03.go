package main

import (
	"sync"
)

func main() {
	var x []int
	s := sync.WaitGroup{}

	s.Add(1)
	go func() {
		defer s.Done()
		x = make([]int, 10)

	}()

	s.Add(1)
	go func() {
		defer s.Done()
		x = make([]int, 1000000)
	}()

	x[999999] = 1

	s.Wait()
}

