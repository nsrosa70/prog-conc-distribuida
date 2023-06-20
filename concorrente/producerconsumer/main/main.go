package main

import (
	"aulas/concorrente/producerconsumer/event"
	monitor "aulas/concorrente/producerconsumer/monitor/impl"
	"fmt"
	"strconv"
	"sync"
	"time"
)

func producer(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	// generate an event
	e := event.Event{E: "event" + strconv.Itoa(id)}

	// publish event
	EB.Add(e)
}

func consumer(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	// consume event
	e := EB.Get()
	e.E = strconv.Itoa(id) + ":" + e.E

	// process event
	e.Process(e)
}

//var EB *mutex.MutexEventBuffer
//var EB *cond.CondEventBuffer
//var EB *channel.ChanEventBuffer
//var EB *semaforo.SemaforoEventBuffer
var EB *monitor.MonitorEventBuffer

func main() {
	wg := sync.WaitGroup{}

	n := 100        // number of producers/consumers
	sample := 10000 // sample size

	//EB = mutex.NewMutexEventBuffer(1)
	//EB  = cond.NewCondEventBuffer(1)
	//EB  = channel.NewChanEventBuffer(1)
	//EB  = semaforo.NewSemaforoEventBuffer(1)
	EB = monitor.NewMonitorEventBuffer(1)

	t1 := time.Now()
	for idx := 0; idx < sample; idx++ {
		for i := 0; i < n; i++ {
			wg.Add(1)
			go consumer(i, &wg)
		}

		//time.Sleep(10 * time.Second)

		for i := 0; i < n; i++ {
			wg.Add(1)
			go producer(i, &wg)
		}
		wg.Wait()
	}
	t2 := time.Now().Sub(t1).Milliseconds()

	fmt.Println(float64(t2) / float64(sample))
}
