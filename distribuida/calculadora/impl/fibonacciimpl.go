package impl

import (
	"aulas/distribuida/calculadora/grpc/fibonacci"
	"context"
)

type Fibonacci struct{}

func (f *Fibonacci) Fibo(ctx context.Context, in *fibonacci.Request) (*fibonacci.Reply, error) {
	return &fibonacci.Reply{N: FibonacciRecursion(in.P1)}, nil
}

func FibonacciRecursion(n int32) int32 {
	if n <= 1 {
		return n
	}
	return FibonacciRecursion(n-1) + FibonacciRecursion(n-2)
}
