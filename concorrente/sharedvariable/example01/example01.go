package main

import (
	"fmt"
	"sync"
)

func f02() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			fmt.Println(i)
		}()
	}

	wg.Wait()
	fmt.Println("Done")
}

func f01() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		x := i

		go func() {
			defer wg.Done()
			fmt.Println(x)
		}()

	}

	wg.Wait()
	fmt.Println("Done")
}


func main(){
	//runtime.GOMAXPROCS(4)
	//go f01()
	go f02()

	fmt.Scanln()
}