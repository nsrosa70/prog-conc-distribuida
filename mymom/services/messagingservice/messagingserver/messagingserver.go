package main

import (
	"fmt"
	"strconv"
	"test/mymom/services/messagingservice/event"
	messaginginvoker "test/mymom/services/messagingservice/invoker"
	messagingproxy "test/mymom/services/messagingservice/proxy"
	"test/shared"
	"time"
)

const Fila = "Fila 1"
const RequestQueue = "RequestQueue"
const ReplyQueue = "ReplyQueue"

func main() {

	go Broker()
	time.Sleep(500 * time.Millisecond)

	go Subscriber()
	go Publisher()

	//go Servidor()
	//go Cliente()

	fmt.Println("'Servidor de Filas' em execução...")
	fmt.Scanln()
}

func Broker() {
	// Start messagingservice invoker
	i := messaginginvoker.New(shared.LocalHost, shared.MessagingPort)
	i.Invoke()

	fmt.Println("Messaging Server finished")
}

func Publisher() {

	// Obtain messagingservice proxy
	messaging := messagingproxy.New(shared.LocalHost, shared.MessagingPort)

	// Publish a message
	for i := 0; i < 1000; i++ {
		e := event.Event{E: "Oi[" + strconv.Itoa(i) + "]"}
		messaging.Publish(Fila, e)
	}
}

func Subscriber() {
	// Obtain messagingservice proxy
	messaging := messagingproxy.New(shared.LocalHost, shared.MessagingPort)

	// Subscriber as consumer
	ch := messaging.Consume(Fila)

	// Consume messages
	for e := range *ch {
		fmt.Println("Subscriber:: ", e.E)
	}
}

func Cliente() {
	// Obtain messagingservice proxy
	messaging := messagingproxy.New(shared.LocalHost, shared.MessagingPort)

	// Subscribe to Reply queue
	ch := messaging.Consume(ReplyQueue)

	e := event.Event{E: "Add(1,2)"}
	messaging.Publish(RequestQueue, e)

	r := <-*ch

	fmt.Println(r)
}

func Servidor() {
	// Obtain messagingservice proxy
	messaging := messagingproxy.New(shared.LocalHost, shared.MessagingPort)

	// Subscribe to Reply queue
	ch := messaging.Consume(RequestQueue)

	req := <-*ch

	fmt.Println(req)

	e := event.Event{E: "Add(1,2)"}
	messaging.Publish(ReplyQueue, e)
}
