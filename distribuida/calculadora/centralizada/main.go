package main

import (
	"distribuida/calculadora/impl"
	"fmt"
	"shared"
	"time"
)

func cliente(ch chan interface{}) {

	// loop
	for i := 0; i < shared.SAMPLE_SIZE; i++ {

		t1 := time.Now()

		// prepara request
		req := shared.Request{Op: "add", P1: i, P2: i}

		// envia request
		ch <- req

		// recebe resposta
		<-ch

		t2 := time.Now()
		x := float64(t2.Sub(t1).Nanoseconds()) / 1000000
		fmt.Println(x)
	}
}

func servidor(ch chan interface{}) {

	for {
		// recebe request
		msgFromClient := <-ch
		req := msgFromClient.(shared.Request)

		// processa request
		r := impl.Calculadora{}.InvocaCalculadora(req)

		// envia resposta
		ch <- r
	}
}

func main() {
	ch := make(chan interface{})

	go servidor(ch)
	go cliente(ch)

	fmt.Scanln()
}
