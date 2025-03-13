package core

import (
	"test/myrpc/distribution/requestor"
)

type Core struct {
	R requestor.Requestor
}

func NewCore(h string, p int) *Core {
	r := new(Core)
	r.R = requestor.NewRequestor(h, p)
	return r
}
