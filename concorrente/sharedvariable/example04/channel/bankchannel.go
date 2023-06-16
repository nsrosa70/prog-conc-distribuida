package channel

import (
	"fmt"
	"sync"
)

var (
	sema    = make(chan int, 1)
	balance int
)

func Deposit(amount int) {
	sema <- 1
	balance = balance + amount
	<-sema
}

func Balance() int {
	sema <- 1
	b := balance
	<-sema
	return b
}

func SetBalance(b int) {
	sema <- 1
	balance = 0
	<-sema
}

func Transaction() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	// Bob:
	go func() {
		defer wg.Done()
		Deposit(200) // A1
		fmt.Println("=", Balance())
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		Deposit(100) // A2
	}()

	wg.Wait()
	return
}

