package main

import (
	"fmt"
	"sync"
	"time"
)

func f1(wg *sync.WaitGroup){
	defer wg.Done()

	for i := 0; i < 10000; i ++{
		fmt.Println("f1",i)
		time.Sleep(1 * time.Millisecond)
	}
}

func f2(wg *sync.WaitGroup){
	defer wg.Done()

	for i := 0; i < 10000; i ++{
		fmt.Println("***** f2 ******",i)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup

	//runtime.GOMAXPROCS(1)

	t1 := time.Now()

	wg.Add(1)
	go f1(&wg)

	wg.Add(1)
	go f2(&wg)

	wg.Wait()

	fmt.Println("Tempo Total: ",time.Now().Sub(t1))
}
