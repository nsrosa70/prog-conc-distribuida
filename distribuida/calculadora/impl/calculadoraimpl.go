package impl

import "aulas/distribuida/calculadora/shared"

type Calculadora struct{}

func (Calculadora) InvocaCalculadora(req shared.Request) int {
	var r int

	op := req.Op
	p1 := req.P1
	p2 := req.P2

	switch op {
	case "add":
		r = Calculadora{}.Add(p1, p2)
	case "sub":
		r = Calculadora{}.Sub(p1, p2)
	case "mul":
		r = Calculadora{}.Mul(p1, p2)
	case "div":
		r = Calculadora{}.Div(p1, p2)
	}
	return r
}

func (Calculadora) Add(x int, y int) int {
	return x + y
}

func (Calculadora) Sub(x int, y int) int {
	return x - y
}

func (Calculadora) Mul(x int, y int) int {
	return x * y
}

func (Calculadora) Div(x int, y int) int {
	return x / y
}
