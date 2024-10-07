package main

import (
	"fmt"
	"test/mymiddleware/distribution/proxies/calculadora"
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

	// Chamada remota a Calculadora
	fmt.Println(calc.Som(1, 2))
}
