package impl

import (
	"context"
	"distribuida/calculadora/grpc/calculadora"
)

type CalculadoraGRPC struct{}

func (s *CalculadoraGRPC) Add(ctx context.Context, in *calculadora.Request) (*calculadora.Reply, error) {
	return &calculadora.Reply{N: in.P1+in.P2}, nil
}

func (s *CalculadoraGRPC) Sub(ctx context.Context, in *calculadora.Request) (*calculadora.Reply, error) {
	return &calculadora.Reply{N: in.P1-in.P2}, nil
}

func (s *CalculadoraGRPC) Mul(ctx context.Context, in *calculadora.Request) (*calculadora.Reply, error) {
	return &calculadora.Reply{N: in.P1*in.P2}, nil
}

func (s *CalculadoraGRPC) Div(ctx context.Context, in *calculadora.Request) (*calculadora.Reply, error) {
	return &calculadora.Reply{N: in.P1/in.P2}, nil
}
