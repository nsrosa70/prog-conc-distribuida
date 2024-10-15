package messagingproxy

import (
	"test/mymom/services/messagingservice/event"
	"test/myrpc/distribution/requestor"
	srh2 "test/myrpc/infrastructure/srh"
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

func (h *MessagingProxy) Consume(_p1 string) *chan string {

	// 1. Configure input parameters
	params := make([]interface{}, 1)
	params[0] = _p1

	// Configure remote request
	req := shared.Request{Op: "Consume", Params: params}

	// Prepare invocation to Requestor
	inv := shared.Invocation{Ior: h.Ior, Request: req}

	// 3. Create call back server first
	ch := make(chan string)
	go Callback(ch)

	// 4. Invoke Requestor
	requestor := requestor.Requestor{}
	_ = requestor.Invoke(inv)

	//4. Return callback channel
	return &ch
}

func Callback(ch chan string) {
	srh := srh2.NewSRH(shared.LocalHost, shared.CallBackPort)
	for {
		msg := srh.Receive()
		ch <- string(msg)
	}
}
