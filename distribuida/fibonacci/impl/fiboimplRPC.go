package impl

type FibonacciRPC struct{}

func (t *FibonacciRPC) Fibo(n *int, reply *int) error {
	*reply = fibonacci(*n)
	return nil
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
