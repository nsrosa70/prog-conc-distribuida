package main

import (
	"fmt"
	proxies "test/mymiddleware/distribution/proxies/calculadora"
	"test/shared"
)

func main() {
	// Obtain proxy
	ior := shared.IOR{Host: "localhost", Port: 1313, Id: 3535, TypeName: "Hello"}
	p := proxies.CalculadoraProxy{Ior: ior}

	for i := 0; i < 10; i++ {
		fmt.Println(i, p.Soma(i, i), p.Diferenca(i, i), p.Multiplicacao(i, i), p.Divisao(i, i))
	}
}
