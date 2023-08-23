package impl

import (
	gen "aulas/distribuida/fibonacci/proto"
	"golang.org/x/net/context"
)

type FibonacciRPC struct{}

func (t *FibonacciRPC) Fibo(ctx context.Context, req *gen.RequestFibo) (*gen.ReplyFibo, error) {
	reply := &gen.ReplyFibo{}
	reply.N = fibonacci(req.P1)
	return reply, nil
}

func fibonacci(n int32) int32 {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
