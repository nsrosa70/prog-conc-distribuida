package impl

type CalculadoraGoRPC struct{}

/*
func (CalculadoraGoRPC) InvocaCalculadora(req shared.Request) int {
	var r int

	op := req.Op
	p1 := req.P1
	p2 := req.P2

	switch op {
	case "add":
		r = CalculadoraGoRPC{}.Add(p1, p2)
	case "sub":
		r = CalculadoraGoRPC{}.Sub(p1, p2)
	case "mul":
		r = CalculadoraGoRPC{}.Mul(p1, p2)
	case "div":
		r = CalculadoraGoRPC{}.Div(p1, p2)
	}
	return r
}
*/

func (CalculadoraGoRPC) Add(x int, y int) int {
	return x + y
}

func (CalculadoraGoRPC) Sub(x int, y int) int {
	return x - y
}

func (CalculadoraGoRPC) Mul(x int, y int) int {
	return x * y
}

func (CalculadoraGoRPC) Div(x int, y int) int {
	return x / y
}
