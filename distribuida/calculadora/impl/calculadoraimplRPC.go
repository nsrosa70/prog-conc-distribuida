package impl

import (
	"aulas/distribuida/calculadora/shared"
	"errors"
)

type CalculadoraRPC struct{}

func (t *CalculadoraRPC) Mul(args *shared.Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}
func (t *CalculadoraRPC) Sub(args *shared.Args, reply *int) error {
	*reply = args.A - args.B
	return nil
}
func (t *CalculadoraRPC) Add(args *shared.Args, reply *int) error {
	*reply = args.A + args.B
	return nil
}
func (t *CalculadoraRPC) Div(args *shared.Args, quo *shared.Quotient) error {
	if args.B == 0 {
		return errors.New("Divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}
