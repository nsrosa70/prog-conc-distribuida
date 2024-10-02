package main

import (
	"fmt"
	"test/mymiddleware/distribution/proxies/calculadora"
	"test/mymiddleware/distribution/proxies/fibonacci"
	namingproxy "test/mymiddleware/services/naming/proxy"
	"test/shared"
)

func main() {
	Cliente()
}

func Cliente() {

	// Obtain proxies
	naming := namingproxy.New(shared.LocalHost, shared.NamingPort)
	calc := calculadoraproxy.New(naming.Find("Calculadora"))
	fibo := fibonacciproxy.New(shared.LocalHost, shared.FibonacciPort)

	// Invoke services
	for i := 0; i < 1000; i++ {
		fmt.Println(i, calc.Som(i, i), calc.Dif(i, i), calc.Mul(i, i), calc.Div(i, i))
		fmt.Println(i, fibo.Fibo(i))
	}
}
