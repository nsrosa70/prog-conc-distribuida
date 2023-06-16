package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//mu := sync.Mutex{}

	cond := sync.NewCond(new(sync.Mutex))
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(3)

	go func() { // G1
		defer waitGroup.Done()

		fmt.Println("G1 vai enviar o broadcast em 10s")
		time.Sleep(10 * time.Second)
		fmt.Println("G1 acorda as outras goroutines")
		cond.Broadcast()
	}() // G1 (envia o broadcast)

	go func() { // G2
		defer waitGroup.Done()
		defer cond.L.Unlock()

		cond.L.Lock()
		fmt.Println("G2 esperando pelo broadcast")
		cond.Wait()
		fmt.Println("G2 recebeu o broadcast")
	}() // G2 (espera pelo broadcast)

	go func() { // G3
		defer waitGroup.Done()
		defer cond.L.Unlock()

		cond.L.Lock()
		fmt.Println("G3 esperando pelo broadcast")
		cond.Wait()
		fmt.Println("G3 recebeu o broadcast")
	}() // G3 (espera pelo broadcast)

	fmt.Println("'main' esperando pela finalização das outras goroutines")
	waitGroup.Wait()
	fmt.Println("'main' finalizou")
}
