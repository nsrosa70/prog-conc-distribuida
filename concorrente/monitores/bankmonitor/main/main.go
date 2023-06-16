package main

import (
	bank "concorrente/monitores/bankmonitor/impl"
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	b := &bank.Bank{}
	b.Init()

	wg.Add(10)
	for i := 0; i < 5; i ++ {

		go func() {
			defer wg.Done()
			b.Deposit(100)
		}()

		go func() {
			defer wg.Done()
			b.Withdraw(100)
		}()
	}
	wg.Wait()

	fmt.Println(b.GetBalance())
}
