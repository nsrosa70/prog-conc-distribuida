package main

import (
	"fmt"
	"os"
	"sharedvariable/example04/channel"
)

func main() {
	for {
		channel.SetBalance(0)
		//mutex.SetBalance(0)
		//rwmutex.SetBalance(0)

		channel.Transaction()
		//mutex.Transaction()
		//rwmutex.Transaction()

		if channel.Balance() < 300 {
		//if channel.Balance() < 300 {
		//	if channel.Balance() < 300 {
			fmt.Println("*************** BINGO ****************", channel.Balance())
			//fmt.Println("*************** BINGO ****************", mutex.Balance())
			//fmt.Println("*************** BINGO ****************", rwmutex.Balance())
			os.Exit(0)
		}
	}
}
