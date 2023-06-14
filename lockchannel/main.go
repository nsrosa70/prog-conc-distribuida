package main

import (
	"fmt"
)

var x int

func f2(ch chan bool) {
	<-ch
	x--
	fmt.Println("f2", x)
	ch <- true
}

func f1(ch chan bool) {
	<-ch
	x++
	fmt.Println("f1", x)
	ch <- true
}

func f3(ch chan bool) {
	<-ch
	x++
	fmt.Println("f3", x)
	ch <- true
}

func main() {
	ch := make(chan bool, 1)

	ch <- true
	go f1(ch)
	go f2(ch)
	go f3(ch)

	fmt.Scanln()
}
