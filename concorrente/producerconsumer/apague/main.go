package main

import (
	"aulas/concorrente/producerconsumer/event"
	"aulas/concorrente/producerconsumer/eventbuffer"
	"fmt"
	"strconv"
	"sync"
)

type CondEventBuffer struct {
	cond *sync.Cond
	eb   eventbuffer.EventBuffer
}

const NumeroDeProdutores = 1
const NumeroDeConsumidores = 1
const CapacidadeDoBuffer = 1000 // 1, 100, 1.000
const NumeroDeItens = 100

var EB eventbuffer.IEventBuffer

func main() {
	wg := sync.WaitGroup{}

	EB = NewCondEventBuffer(CapacidadeDoBuffer, NumeroDeProdutores*NumeroDeItens) // tipo de primitiva
	for i := 0; i < NumeroDeConsumidores; i++ {                                   // inicia os consumidores
		wg.Add(1)
		go consumidor(i, &wg)
	}

	for i := 0; i < NumeroDeProdutores; i++ { // inicia os produtores
		wg.Add(1)
		go produtor(i, &wg)
	}

	wg.Wait()
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
		if e.E == "" {
			fmt.Println("No more itens") // não há mais eventos
			break
		}
	}
}

func NewCondEventBuffer(capacity, consumed int) *CondEventBuffer {
	eventBuffer := eventbuffer.EventBuffer{Capacity: capacity, Buffer: []event.Event{}, Consumed: consumed}
	return &CondEventBuffer{
		cond: sync.NewCond(&sync.Mutex{}),
		eb:   eventBuffer,
	}
}

func (b *CondEventBuffer) Add(e event.Event) {
	b.cond.L.Lock()
	for len(b.eb.Buffer) == b.eb.Capacity {
		b.cond.Wait()
	}
	b.eb.Buffer = append(b.eb.Buffer, e)
	b.cond.Broadcast()
	b.cond.L.Unlock()
}

func (b *CondEventBuffer) Get() event.Event {
	b.cond.L.Lock()
	for len(b.eb.Buffer) == 0 {
		if b.eb.Consumed == 0 {
			b.cond.L.Unlock()
			return event.Event{E: ""}
		}
		b.cond.Wait()
	}
	e := b.eb.Buffer[0]
	b.eb.Buffer = b.eb.Buffer[1:]
	b.eb.Consumed--
	b.cond.Broadcast()
	b.cond.L.Unlock()

	return e
}
