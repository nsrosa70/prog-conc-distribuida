package calculadora

type Calculadora struct{}

func (Calculadora) Soma(p1, p2 int) int {
	return p1 + p2
}

func (Calculadora) Diferenca(p1, p2 int) int {
	return p1 - p2
}

func (Calculadora) Multiplicacao(p1, p2 int) int {
	return p1 * p2
}

func (Calculadora) Divisao(p1, p2 int) int {
	if p2 == 0 {
		return 0
	} else {
		return p1 / p2
	}
}
