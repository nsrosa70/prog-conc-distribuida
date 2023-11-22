package main

import (
	"aulas/concorrente/producerconsumer/event"
	"aulas/concorrente/producerconsumer/eventbuffer"
	"aulas/concorrente/producerconsumer/shared"
	"fmt"
	"strconv"
	"sync"
	"time"
)

const NumeroDeProdutores = 100
const NumeroDeConsumidores = 100
const TamanhoDaAmostra = 100
const CapacidadeDoBuffer = 1 // 1, 100, 1.000
const NumeroDeItens = 10000

var EB eventbuffer.IEventBuffer

func main() {
	wg := sync.WaitGroup{}

	for idx := 0; idx < TamanhoDaAmostra; idx++ {
		EB = shared.NewEventBuffer("Channel", CapacidadeDoBuffer, NumeroDeProdutores*NumeroDeItens) // tipo de primitiva
		t1 := time.Now()
		for i := 0; i < NumeroDeConsumidores; i++ { // inicia os consumidores
			wg.Add(1)
			go consumidor(i, &wg)
		}

		for i := 0; i < NumeroDeProdutores; i++ { // inicia os produtores
			wg.Add(1)
			go produtor(i, &wg)
		}

		wg.Wait()

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
		e := EB.Get()
		e.Process(e)
		if e.E == "" { // não há mais eventos
			break
		}
	}
}
