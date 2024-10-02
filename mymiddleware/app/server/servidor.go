package main

import (
	"fmt"
	"test/mymiddleware/distribution/invokers/calculadora"
	"test/mymiddleware/distribution/invokers/fibonacci"
)

func main() {
	// Start calculadora invoker
	calc := calculadora.Invoker{}

	// Start fibonacci invoker
	fibo := fibonacci.Invoker{}

	go calc.Invoke("localhost", 1313)
	go fibo.Invoke("localhost", 1314)

	fmt.Println("Servidor pronto...")
	fmt.Scanln()
}
