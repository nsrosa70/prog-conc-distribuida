package main

import (
	"fmt"
)

func g1(ch chan int) {
	n := 1
	fmt.Println("g1 sends", n)
	ch <- 1 // send int through channel
}

func g2(ch chan int) {
	n := <-ch // receive int from channel
	fmt.Println("g2 receives", n)
}

func main() {
	ch := make(chan int) // create a channel

	go g1(ch) // start g1
	go g2(ch) // start g2

	_, _ = fmt.Scanln()
}
