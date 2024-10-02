package main

import (
	"fmt"
	calculadorainvoker "test/mymiddleware/distribution/invokers/calculadora"
	calculadoraproxy "test/mymiddleware/distribution/proxies/calculadora"
	naminginvoker "test/mymiddleware/services/naming/invoker"
	namingproxy "test/mymiddleware/services/naming/proxy"
	"test/shared"
	"time"
)

func main() {

	go namingServer()
	time.Sleep(100 * time.Millisecond)
	go calcServer()
	time.Sleep(100 * time.Millisecond)
	go calcClient()

	fmt.Println("Servidor de Nomes em execução...")
	fmt.Scanln()
}

func namingServer() {
	// Start naming invoker
	i := naminginvoker.New(shared.LocalHost, shared.NamingPort)
	go i.Invoke()
}

func calcServer() {
	// Create proxy of naming server
	naming := namingproxy.New(shared.LocalHost, shared.NamingPort)

	// Create invoker of calculadora
	calcInv := calculadorainvoker.New(shared.LocalHost, shared.CalculadoraPort)

	// Register Calculadora in Naming server
	calcIor := shared.NewIOR(calcInv.Ior.Host, calcInv.Ior.Port)
	naming.Bind("Calculadora", calcIor)

	// Invoke Calculadora Invoker
	calcInv.Invoke()
}

func calcClient() {
	naming := namingproxy.New(shared.LocalHost, shared.NamingPort)
	calc := calculadoraproxy.New(naming.Find("Calculadora"))

	for i := 0; i < 100; i++ {
		fmt.Println(calc.Som(i, i))
	}
	fmt.Scanln()
}
