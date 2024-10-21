package requestor

import (
	"test/myrpc/distribution/marshaller"
	"test/myrpc/distribution/miop"
	"test/myrpc/infrastructure/crh"
	"test/shared"
)

type Requestor struct {
}

func NewRequestor() Requestor {
	return Requestor{}
}

func (Requestor) Invoke(i shared.Invocation) shared.Termination {
	// 1. Create MIOP packet
	miopReqPacket := miop.CreateRequestMIOP(i.Request.Op, i.Request.Params)

	// 2. Serialise MIOP packet
	m := marshaller.Marshaller{}
	b := m.Marshall(miopReqPacket)

	// 3. Create & invoke CRH
	c := crh.NewCRH(i.Ior.Host, i.Ior.Port)
	r := c.SendReceive(b)

	// 4. Extract reply from subscriber
	miopRepPacket := m.Unmarshall(r)
	rt := miop.ExtractReply(miopRepPacket)

	t := shared.Termination{Rep: rt}

	return t
}
