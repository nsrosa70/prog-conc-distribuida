package core

import (
	"test/myrpc/distribution/marshaller"
	"test/myrpc/distribution/requestor"
	"test/myrpc/infrastructure/crh"
)

type Core struct {
	_Requestor  requestor.Requestor
	_Crh        crh.CRH
	_Marshaller marshaller.Marshaller
}

func NewCore() Core {
	r := Core{}
	/*	r._Requestor = NewRequestor()
		r._Crh = crh.NewCRH()
		r._Marshaller = marshaller.Marshaller{}
	*/
	return r
}
