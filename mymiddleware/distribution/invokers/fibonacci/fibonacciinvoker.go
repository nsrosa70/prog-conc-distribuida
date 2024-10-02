package fibonacci

import (
	"log"
	"test/mymiddleware/app/businesses/fibonacci"
	"test/mymiddleware/distribution/marshaller"
	"test/mymiddleware/distribution/miop"
	"test/mymiddleware/infrastructure/srh"
)

type Invoker struct{}

func (Invoker) Invoke(h string, p int) {
	s := srh.NewSRH(h, p)
	m := marshaller.Marshaller{}
	c := fibonacci.Fibonacci{}
	miopPacket := miop.Packet{}
	var rep int

	for {
		// Invoke SRH
		b := s.Receive()

		// Unmarshall miop packet
		miopPacket = m.Unmarshall(b)

		// Extract request from client
		r := miop.ExtractRequest(miopPacket)

		_p1 := int(r.Params[0].(float64))

		// Demultiplex request
		switch r.Op {
		case "Fibo":
			rep = c.Fibo(_p1)
		default:
			log.Fatal("Operation unknown:: " + r.Op)
		}

		// Prepare reply
		var params []interface{}
		params = append(params, rep)

		// Create miop reply packet
		miop := miop.CreateReplyMIOP(params)

		// Marshall miop packet
		b = m.Marshall(miop)

		// Send marshalled packet
		s.Send(b)
	}
}
