package fibonacci

type Fibonacci struct{}

func (Fibonacci) Fibo(p1 int) int {
	var n2, n1 = 0, 1

	if p1 <= 1 {
		return p1
	}
	for i := 2; i <= p1; i++ {
		n2, n1 = n1, n1+n2
	}
	return n1
}
