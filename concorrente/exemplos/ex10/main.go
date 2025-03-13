package main

import (
	"fmt"
	"sync"
)

func FazQuaseNada(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Quase Nada")
}

func FazTiquinho(wg *sync.WaitGroup) {
	//defer wg.Done()
	fmt.Println("Tiquinho")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go FazQuaseNada(&wg)

	wg.Add(1)
	go FazTiquinho(&wg)

	fmt.Println("Main")

	wg.Wait()
}
