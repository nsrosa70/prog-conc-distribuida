package main

import (
	"fmt"
)

func main05() {
	ch01 := make(chan string)
	ch02 := make(chan string)

	go func(ch chan string) {
		ch <- "Função 01"
	}(ch01) // Função 01

	go func(ch chan string) {
		ch <- "Função 02"
	}(ch02) // Função 02

	go func(ch01,ch02 chan string) {
		select {
		case m := <-ch01:
			fmt.Println(m)
		case m := <-ch02:
			fmt.Println(m)
		default:
			fmt.Println("Nada recebido")
		}
	}(ch01,ch02)

	//x := "teste"

	//fmt.Println("O tipo da variavel 'x' é",reflect.TypeOf(x).String())

	fmt.Scanln()
}


func main() {
	ch := make(chan int,2)

	ch <- 1
	ch <- 2

	for m := range ch {
		fmt.Println(m)
	}
}

func main01() {
	ch := make(chan int, 2)

	ch <- 1
	close(ch)
	_, isClosed := <-ch

	if isClosed {
		fmt.Println("Canal 'ch' está fechado")
	} else {
		fmt.Println(<-ch)
	}
}

func main02() {
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
