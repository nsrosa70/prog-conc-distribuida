package shared

import (
	"fmt"
	"os"
	"test/concorrente/producerconsumer/channel"
	"test/concorrente/producerconsumer/cond"
	"test/concorrente/producerconsumer/eventbuffer"
	monitor "test/concorrente/producerconsumer/monitor/impl"
	"test/concorrente/producerconsumer/mx"
	"test/concorrente/producerconsumer/semaforo"
)

func NewEventBuffer(t string, capacity, consumed int) eventbuffer.IEventBuffer {
	switch t {
	case "Channel":
		eb := channel.NewChanEventBuffer(capacity, consumed)
		return eb
	case "Cond":
		eb := cond.NewCondEventBuffer(capacity, consumed)
		return eb
	case "Mutex":
		eb := mx.NewMutexEventBuffer(capacity, consumed)
		return eb
	case "Semaforo":
		eb := semaforo.NewSemaforoEventBuffer(capacity, consumed)
		return eb
	case "Monitor":
		eb := monitor.NewMonitorEventBuffer(capacity, consumed)
		return eb
	default:
		fmt.Println("Tipo de primitiva", t, "é desconhecida!")
		os.Exit(0)
	}
	return *new(eventbuffer.IEventBuffer)
}
