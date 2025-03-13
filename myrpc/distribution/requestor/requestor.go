package requestor

import (
	"test/myrpc/distribution/marshaller"
	"test/myrpc/distribution/miop"
	"test/myrpc/infrastructure/crh"
	"test/shared"
)

type Requestor struct {
	M marshaller.Marshaller
	C crh.CRH
}

func NewRequestor(h string, p int) Requestor {
	c := crh.NewCRH(h, p)
	m := marshaller.Marshaller{}
	r := Requestor{M: m, C: *c}

	return r
}

func (req *Requestor) Invoke(i shared.Invocation) shared.Termination {
	// 1. Create MIOP packet
	miopReqPacket := miop.CreateRequestMIOP(i.Request.Op, i.Request.Params)

	// 2. Serialise MIOP packet
	//m := marshaller.Marshaller{}
	//b := m.Marshall(miopReqPacket)
	b := req.M.Marshall(miopReqPacket)

	// 3. Create & invoke CRH
	//c := crh.NewCRH(i.Ior.Host, i.Ior.Port)
	r := req.C.SendReceive(b)

	// 4. Extract reply from subscriber
	//miopRepPacket := m.Unmarshall(r)
	miopRepPacket := req.M.Unmarshall(r)
	rt := miop.ExtractReply(miopRepPacket)

	t := shared.Termination{Rep: rt}

	return t
}
