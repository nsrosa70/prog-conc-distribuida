package impl

import (
	gen "aulas/distribuida/messagingservice/grpc/proto"
	"golang.org/x/net/context"
)

type CalculadoraRPC struct {
}

func (t *CalculadoraRPC) Add(ctx context.Context, request *gen.Request) (*gen.Reply, error) {
	reply := &gen.Reply{}
	reply.N = request.P1 + request.P2
	return reply, nil
}
func (t *CalculadoraRPC) Sub(ctx context.Context, request *gen.Request) (*gen.Reply, error) {
	reply := &gen.Reply{}
	reply.N = request.P1 - request.P2
	return reply, nil
}
func (t *CalculadoraRPC) Mul(ctx context.Context, request *gen.Request) (*gen.Reply, error) {
	reply := &gen.Reply{}
	reply.N = request.P1 * request.P2
	return reply, nil
}
func (t *CalculadoraRPC) Div(ctx context.Context, request *gen.Request) (*gen.Reply, error) {
	reply := &gen.Reply{}
	reply.N = request.P1 / request.P2
	return reply, nil
}
