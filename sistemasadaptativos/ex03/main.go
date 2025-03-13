package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const MonitorTime = 1 // s
const ContActive = true
const Min = 1   // msg/s
const Max = 100 // msgs/s
const Goal = 50
const SampleSize = 100

func Environment(chToGerenciador chan int) {
	for {
		chToGerenciador <- myRandon(Min, Max) //rate msgs/s
		time.Sleep(MonitorTime * time.Second)
	}
}

func SistemaGerenciador(chFromEnv, chToGerenciado chan int) {
	c := Initialise(Goal, Min, Max)
	s := 0
	for i := 0; i < SampleSize; i++ {
		m := <-chFromEnv
		if ContActive {
			s = c.Update(m)
		} else {
			s = m
		}
		fmt.Printf("%v\n", s)
		chToGerenciado <- s
	}
	os.Exit(1)
}

func SistemaGerenciado(chFromGerenciador chan int) {
	r := <-chFromGerenciador
	for {
		select {
		case r = <-chFromGerenciador:
		default:
		}
		//fmt.Print("a")
		//fmt.Printf("[%v]", r)
		//fmt.Printf("%.2f\n", float64(1000)/float64(s))
		time.Sleep(time.Duration(1/r) * time.Second)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Environment(ch1)
	go SistemaGerenciador(ch1, ch2)
	go SistemaGerenciado(ch2)

	fmt.Scanln()
}

type OnOff struct {
	Max  int
	Min  int
	Goal int
}

func Initialise(g, min, max int) OnOff {
	c := OnOff{Max: max, Min: min, Goal: g}
	return c
}

func (c OnOff) Update(m int) int {
	if m > c.Goal {
		return c.Min
	} else {
		return c.Max
	}
}

func myRandon(min, max int) int {
	var r int
	for {
		r = rand.Intn(max)
		if r >= min && r <= max {
			break
		}
	}
	return r
}
