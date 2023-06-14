package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func g1(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i < 1000; i++ {
		fmt.Println("g1")
	}
}

func g2(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i < 1000; i++ {
		fmt.Println("g2")
	}
}

func main() {
	wg := sync.WaitGroup{}
	//x := runtime.GOMAXPROCS(1)

	runtime.GOMAXPROCS(1)
	//fmt.Println(runtime.NumCPU())

	//fmt.Println(x, runtime.GOMAXPROCS(3))

	wg.Add(2)
	t1 := time.Now()
	go g1(&wg)
	go g2(&wg)

	//go g2()

	wg.Wait()

	fmt.Println(time.Now().Sub(t1))

}
