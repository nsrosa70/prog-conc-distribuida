package main

import (
	"fmt"
	"time"
)

func f1(ch chan string) {
	time.Sleep(1 * time.Nanosecond)
	ch <- "from f1"
}

func f2(ch chan string) {
	time.Sleep(1 * time.Nanosecond)
	ch <- "from f2"
}

func f3(ch1, ch2 chan string) {
	msg := ""
	select {
	case msg = <-ch1:
		fmt.Println(msg)
	case msg = <-ch2:
		fmt.Println(msg)
	default:
		fmt.Println("*********** No message ***************")
		return
	}
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	end := make(chan bool, 1)

	go f1(ch1)
	go f2(ch2)
	go f3(ch1, ch2)
	//fmt.Scanln()

	end <- true
}
