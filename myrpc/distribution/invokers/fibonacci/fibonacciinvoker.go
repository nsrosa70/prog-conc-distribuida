package fibonacciinvoker

import (
	"log"
	"test/myrpc/app/businesses/fibonacci"
	"test/myrpc/distribution/marshaller"
	"test/myrpc/distribution/miop"
	"test/myrpc/infrastructure/srh"
	"test/shared"
)

type FibonacciInvoker struct {
	Ior shared.IOR
}

func New(h string, p int) FibonacciInvoker {
	ior := shared.IOR{Host: h, Port: p}
	inv := FibonacciInvoker{Ior: ior}

	return inv
}
func (i FibonacciInvoker) Invoke() {
	s := srh.NewSRH(i.Ior.Host, i.Ior.Port)
	m := marshaller.Marshaller{}
	c := fibonacci.Fibonacci{}
	miopPacket := miop.Packet{}
	var rep int

	for {
		// Invoke SRH
		b := s.Receive()

		// Unmarshall miop packet
		miopPacket = m.Unmarshall(b)

		// Extract request from publisher
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
