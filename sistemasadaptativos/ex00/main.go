package main

import (
	"aulas/sistemasadaptativos/shared"
	"fmt"
	"math/rand"
	"time"
)

type Managing struct{}
type Managed struct{}
type SelfAdaptive struct{}

func (Managing) Run(ch chan int) {
	for {
		time.Sleep(shared.MonitorTime * time.Second)
		ch <- rand.Intn(shared.MaxIntervalTime) // time between prints
	}
}

func (Managed) Run(ch chan int) {
	t := 100 // 100 ms
	for {
		select {
		case t = <-ch: // received from managing
		default:
			fmt.Printf("x")
		}
		time.Sleep(time.Duration(t) * time.Millisecond)
	}
}

func (SelfAdaptive) Run() {
	ch := make(chan int)

	go Managing{}.Run(ch)
	go Managed{}.Run(ch)
}

func main() {
	SelfAdaptive{}.Run()

	fmt.Scanln()
}
