package main

import (
	"fmt"
	"sync"
	"time"
)

func G(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("G%d iniciando \n", id)
	time.Sleep(time.Second)
	fmt.Printf("G%d concluida \n", id)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(5)
	for i := 1; i <= 5; i++ {
		go G(i, &wg)
	}
	wg.Wait()

	//fmt.Scanln()

	fmt.Printf("*** Goroutines concluidas ***\n")
}
