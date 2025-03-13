package main

import (
	"fmt"
	"runtime"
	"sync"
	"test/concorrente/synccond/condqueue"
	"time"
)

const TamanhoDaAmostra = 1000
const NumeroDeThreads = 100000
const CapacidadeDaFila = 100
const GoMaxProcs = 6

func main() {
	runtime.GOMAXPROCS(GoMaxProcs)
	for i := 0; i < TamanhoDaAmostra; i++ {
		t1 := time.Now()
		expQueue()
		t2 := time.Now()
		fmt.Println(t2.Sub(t1).Milliseconds())
	}
}

func expQueue() {
	var wg = new(sync.WaitGroup)
	var q = condqueue.NewCondQueue(CapacidadeDaFila)

	for i := 0; i <= NumeroDeThreads; i++ {
		wg.Add(1)
		go func() {
			q.Push(i, wg)
		}()
		wg.Add(1)

		go func() {
			q.Pop(wg)
		}()
	}
}

//var q = mutexqueue.NewMutexQueue(CapacidadeDaFila)
