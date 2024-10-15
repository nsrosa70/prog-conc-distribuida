package messaginginvoker

import (
	"log"
	"test/mymom/services/messagingservice"
	"test/mymom/services/messagingservice/event"
	"test/myrpc/distribution/marshaller"
	"test/myrpc/distribution/miop"
	"test/myrpc/infrastructure/srh"
	"test/shared"
)

type Invoker struct {
	Ior shared.IOR
}

func New(h string, p int) Invoker {
	i := shared.IOR{Host: h, Port: p}
	r := Invoker{Ior: i}
	return r
}

func (i Invoker) Invoke() {
	s := srh.NewSRH(i.Ior.Host, i.Ior.Port)
	m := marshaller.Marshaller{}
	miopPacket := miop.Packet{}

	//var rep interface{}

	// Create an instance of Messaging
	ms := messagingservice.NewMessagingService()

	for {
		// Invoke SRH
		b := s.Receive()

		// Unmarshall miop packet
		miopPacket = m.Unmarshall(b)

		// Extract request from publisher
		r := miop.ExtractRequest(miopPacket)

		go handlingInvocation(r, &ms, *s)
	}
}

func handlingInvocation(r shared.Request, ms *messagingservice.MessagingService, srh srh.SRH) {
	var rep interface{}
	m := marshaller.Marshaller{}

	// Demultiplex request
	switch r.Op {
	case "Publish":
		_p1 := r.Params[0].(string)
		_p2 := r.Params[1].(map[string]interface{})
		_p3 := event.Event{E: _p2["E"].(string)}
		rep = ms.Publish(_p1, _p3)
	case "Consume":
		_p1 := r.Params[0].(string)
		go ms.Consume(_p1) // Callback
		rep = ""           // return nothing
	default:
		log.Fatal("Invoker:: Operation '" + r.Op + "' is unknown:: ")
	}

	// Prepare reply
	var params []interface{}
	params = append(params, rep)

	// Create miop reply packet
	miop := miop.CreateReplyMIOP(params)

	// Marshall miop packet
	b := m.Marshall(miop)

	// Send marshalled packet
	srh.Send(b)
}
