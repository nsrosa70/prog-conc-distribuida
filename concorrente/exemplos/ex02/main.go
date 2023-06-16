package main

import (
	"fmt"
	"time"
)

func Sender(ch chan int) {
	for n := 0; n < 10; n++ {
		ch <- n
		fmt.Println("Send:: ", n)
		time.Sleep(1 * time.Nanosecond)
	}
}

func Receiver(ch chan int) {
	for {
		n := <-ch
		fmt.Println("Receive: ", n)
		time.Sleep(1 * time.Nanosecond)
	}
}

func main() {
	ch := make(chan int,0)

	go Sender(ch)
	go Receiver(ch)

	fmt.Scanln()
}
