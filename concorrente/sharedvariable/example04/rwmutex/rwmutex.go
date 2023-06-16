package rwmutex

import (
	"fmt"
	"sync"
)

var (
	mu      sync.RWMutex
	balance int
)

func Balance() int {
	mu.RLock() // readers lock
	defer mu.RUnlock()
	return balance
}

func Deposit(amount int) {
	mu.Lock()
	balance = balance + amount
	mu.Unlock()
}

func SetBalance(b int) {
	mu.Lock()
	balance = 0
	mu.Unlock()
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
		Deposit(100) // A1
	}()

	wg.Wait()
	return
}