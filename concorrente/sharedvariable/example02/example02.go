package main

import (
	"fmt"
	"os"
	"sharedvariable/example02/bank"
	"sync"
)

func transaction() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	// Bob:
	go func() {
		defer wg.Done()
		bank.Deposit(200) // A1
		fmt.Println("=",bank.Balance())
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		bank.Deposit(100) // A2
	}()

	wg.Wait()
	return
}

func main() {
	for {
		bank.SetBalance(0)

		transaction()

		if bank.Balance() < 300 {
			fmt.Println("*************** BINGO ****************",bank.Balance())
			os.Exit(0)
		}
	}
}
