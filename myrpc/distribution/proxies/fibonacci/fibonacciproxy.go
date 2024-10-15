package fibonacciproxy

import (
	"test/myrpc/distribution/requestor"
	"test/shared"
)

type FibonacciProxy struct {
	Ior shared.IOR
}

func New(h string, p int) FibonacciProxy {
	i := shared.IOR{Host: h, Port: p}
	r := FibonacciProxy{Ior: i}
	return r
}

func (p *FibonacciProxy) Fibo(_p1 int) int {

	// 1. Configure input parameters
	params := make([]interface{}, 1)
	params[0] = _p1

	// Configure remote request
	req := shared.Request{Op: "Fibo", Params: params}

	// Prepare invocation to Requestor
	inv := shared.Invocation{Ior: p.Ior, Request: req}

	// 3. Invoke Requestor
	requestor := requestor.Requestor{}
	r := requestor.Invoke(inv)

	//4. Return something to the publisher
	return int(r.Rep.Result[0].(float64)) // TODO
}
