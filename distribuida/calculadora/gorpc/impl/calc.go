package impl

type Calculadora struct{}

type Request struct {
	P1 int
	P2 int
}

type Reply struct {
	R int
}

func (m *Calculadora) Add(req Request, res *Reply) error {
	res.R = req.P1 + req.P2
	return nil
}

func (m *Calculadora) Sub(req Request, res *Reply) error {
	res.R = req.P1 - req.P2
	return nil
}

func (m *Calculadora) Mul(req Request, res *Reply) error {
	res.R = req.P1 * req.P2
	return nil
}

func (m *Calculadora) Div(req Request, res *Reply) error {
	res.R = req.P1 / req.P2
	return nil
}
