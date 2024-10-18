package messagingproxy

import (
	"encoding/json"
	"log"
	"test/mymom/services/messagingservice/event"
	"test/myrpc/distribution/requestor"
	"test/myrpc/infrastructure/srh"
	"test/shared"
)

type MessagingProxy struct {
	Ior shared.IOR
}

func New(h string, p int) MessagingProxy {
	i := shared.IOR{Host: h, Port: p}
	r := MessagingProxy{Ior: i}
	return r
}

func (h *MessagingProxy) Publish(_p1 string, _p2 event.Event) bool {

	// 1. Configure input parameters
	params := make([]interface{}, 2)
	params[0] = _p1
	params[1] = _p2

	// Configure remote request
	req := shared.Request{Op: "Publish", Params: params}

	// Prepare invocation to Requestor
	inv := shared.Invocation{Ior: h.Ior, Request: req}

	// 3. Invoke Requestor
	requestor := requestor.Requestor{}
	r := requestor.Invoke(inv)

	//4. Return something to the publisher
	return r.Rep.Result[0].(bool)
}

func (h *MessagingProxy) Consume(_p1 string) *chan event.Event {

	// 1.Configure input parameters
	params := make([]interface{}, 1)
	params[0] = _p1

	// 2.Configure remote request
	req := shared.Request{Op: "Consume", Params: params}

	// 3.Prepare invocation to Requestor
	inv := shared.Invocation{Ior: h.Ior, Request: req}

	// 4.Create & invoke callback
	ch := make(chan event.Event)
	go Callback(ch)

	// 5.Invoke Requestor
	requestor := requestor.Requestor{}
	_ = requestor.Invoke(inv)

	//6.Return callback channel
	return &ch
}

func Callback(ch chan event.Event) {
	// Create server (SRH)
	s := srh.NewSRH(shared.LocalHost, shared.CallBackPort)

	// Receive events from broker
	for {
		msg := s.Receive()
		ev := event.Event{}
		err := json.Unmarshal(msg, &ev)
		if err != nil {
			log.Fatal("Messaging:: Callback:: encode error:", err)
		}

		// Send event to Subscriber
		ch <- ev
	}
}
