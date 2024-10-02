package main

import (
	"fmt"
	"test/mymiddleware/distribution/proxies/calculadora"
	"test/mymiddleware/distribution/proxies/fibonacci"
	"test/shared"
)

func main() {
	Cliente()
}

func Cliente() {
	// Obtain proxy
	iorCalc := shared.IOR{Host: "localhost", Port: 1313, Id: 3535, TypeName: "Calculadora"}
	iorFibo := shared.IOR{Host: "localhost", Port: 1314, Id: 3535, TypeName: "Fibonacci"}
	calc := calculadora.CalculadoraProxy{Ior: iorCalc}
	fibo := fibonacci.FibonacciProxy{Ior: iorFibo}

	for i := 0; i < 1000; i++ {
		fmt.Println(i, calc.Som(i, i), calc.Dif(i, i), calc.Mul(i, i), calc.Div(i, i))
		fmt.Println(i, fibo.Fibo(i))
	}
}
