package main

import "fmt"

func f1(ch chan int) {
	fmt.Println("From other function: ", <-ch)
}

func f2(ch chan int) {
	ch <- 1
}

func main() {
	ch := make(chan int)

	go f1(ch)
	go f2(ch)

	fmt.Scanln()
}
