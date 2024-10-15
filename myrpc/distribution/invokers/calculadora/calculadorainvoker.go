package calculadorainvoker

import (
	"log"
	"test/myrpc/app/businesses/calculadora"
	qosobserver "test/myrpc/distribution/interceptors/qos"
	"test/myrpc/distribution/marshaller"
	"test/myrpc/distribution/miop"
	"test/myrpc/infrastructure/srh"
	"test/shared"
)

type CalculadoraInvoker struct {
	Ior shared.IOR
}

func New(h string, p int) CalculadoraInvoker {
	ior := shared.IOR{Host: h, Port: p}
	inv := CalculadoraInvoker{Ior: ior}

	return inv
}

func (i CalculadoraInvoker) Invoke() {
	s := srh.NewSRH(i.Ior.Host, i.Ior.Port)
	m := marshaller.Marshaller{}
	miopPacket := miop.Packet{}
	var rep int

	// Create an instance of Calculadora - Static Instance
	c := calculadora.Calculadora{}

	for {
		// Invoke SRH
		b := s.Receive()

		// Unmarshall miop packet
		miopPacket = m.Unmarshall(b)

		// Extract request from publisher
		r := miop.ExtractRequest(miopPacket)

		_p1 := int(r.Params[0].(float64))
		_p2 := int(r.Params[1].(float64))

		// Demultiplex request & invoke QoS Observer
		qosObserver := qosobserver.QoSObserver{}
		qosObserver.StartTime()
		switch r.Op {
		case "Som":
			rep = c.Som(_p1, _p2)
		case "Dif":
			rep = c.Dif(_p1, _p2)
		case "Mul":
			rep = c.Mul(_p1, _p2)
		case "Div":
			rep = c.Div(_p1, _p2)
		default:
			log.Fatal("Invoker:: Operation '" + r.Op + "' is unknown:: ")
		}
		qosObserver.StopTime()

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
