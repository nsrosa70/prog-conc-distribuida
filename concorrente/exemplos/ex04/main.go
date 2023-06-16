package main

import (
	"sync"
	"sync/atomic"
)

func add(w *sync.WaitGroup, num *int) {
	defer w.Done()
	*num = *num + 1
}

func addAtomic(w *sync.WaitGroup, num *int32) {
	defer w.Done()
	atomic.AddInt32(num, 1)
}

func main() {
	var n int32 = 0 // atomic
	//var m int = 0

	var wg = new(sync.WaitGroup)
	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go addAtomic(wg, &n)  // atomic
		//go add(wg, &m)
	}
	wg.Wait()

	println(n)  // atomic
	//println(m)
}
