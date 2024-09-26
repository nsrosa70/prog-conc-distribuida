package calculadorainvoker

import (
	"log"
	"test/mymiddleware/app/calculadora"
	"test/mymiddleware/distribution/marshaller"
	"test/mymiddleware/distribution/miop"
	"test/mymiddleware/infrastructure/srh"
)

type Invoker struct{}

func (Invoker) Invoke(h string, p int) {
	s := srh.NewSRH(h, p)
	m := marshaller.Marshaller{}
	c := calculadora.Calculadora{}
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
		_p2 := int(r.Params[1].(float64))

		// Demultiplex request
		switch r.Op {
		case "Soma":
			rep = c.Soma(_p1, _p2)
		case "Diferenca":
			rep = c.Diferenca(_p1, _p2)
		case "Multiplicacao":
			rep = c.Multiplicacao(_p1, _p2)
		case "Divisao":
			rep = c.Divisao(_p1, _p2)
		default:
			log.Fatal("Operation unknown")
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
