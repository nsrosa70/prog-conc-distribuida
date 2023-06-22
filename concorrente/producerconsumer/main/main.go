package main

import (
	"aulas/concorrente/producerconsumer/channel"
	"aulas/concorrente/producerconsumer/event"
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

const NumeroDeProdutores = 1
const NumeroDeConsumidores = 2
const TamanhoDaAmostra = 1
const CapacidadeDoBuffer = 1 // 1, 100, 1.000
const NumeroDeItens = 2

var DoneP int32

//var EB *mx.MutexEventBuffer

//var EB *cond.CondEventBuffer

var EB *channel.ChanEventBuffer

//var EB *semaforo.SemaforoEventBuffer

//var EB *monitor.MonitorEventBuffer

//var EB = *new(interface{})

func main() {
	wgP := sync.WaitGroup{}
	wgC := sync.WaitGroup{}

	//EB = primitive("Mutex")
	//EB = mx.NewMutexEventBuffer(CapacidadeDoBuffer)
	//EB = cond.NewCondEventBuffer(CapacidadeDoBuffer)
	EB = channel.NewChanEventBuffer(CapacidadeDoBuffer)
	//EB = semaforo.NewSemaforoEventBuffer(CapacidadeDoBuffer)
	//EB = monitor.NewMonitorEventBuffer(CapacidadeDoBuffer)

	for idx := 0; idx < TamanhoDaAmostra; idx++ {
		t1 := time.Now()
		for i := 0; i < NumeroDeConsumidores; i++ { // inicia os consumidores
			wgC.Add(1)
			go consumidor(i, &wgC)
		}

		for i := 0; i < NumeroDeProdutores; i++ { // inicia os produtores
			wgP.Add(1)
			go produtor(i, &wgP)
		}

		wgP.Wait()
		atomic.StoreInt32(&DoneP, 1)
		wgC.Wait()

		t2 := time.Now().Sub(t1).Milliseconds()
		fmt.Println(t2)
	}
}

func produtor(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < NumeroDeItens; i++ {
		e := event.Event{E: "event [" + strconv.Itoa(id) + "," + strconv.Itoa(i) + "]"} // gera um evento
		EB.Add(e)                                                                       // publica o evento
	}
}

func consumidor(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		if EB.Size() > 0 {
			e := EB.Get()
			//fmt.Println("here")
			e.Process(e)
		} else if atomic.LoadInt32(&DoneP) == 1 {
			break
		}
	}
}
