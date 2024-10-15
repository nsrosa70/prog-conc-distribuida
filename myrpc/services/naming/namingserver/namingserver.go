package main

import (
	"fmt"
	calculadorainvoker "test/myrpc/distribution/invokers/calculadora"
	calculadoraproxy "test/myrpc/distribution/proxies/calculadora"
	naminginvoker "test/myrpc/services/naming/invoker"
	namingproxy "test/myrpc/services/naming/proxy"
	"test/shared"
)

func main() {

	go namingServer()
	//	time.Sleep(100 * time.Millisecond)
	//	go calcServer()
	//	time.Sleep(100 * time.Millisecond)
	//	go calcClient()

	fmt.Println("'Servidor de Nomes' em execução...")
	fmt.Scanln()
}

func namingServer() {
	// Start messagingservice invoker
	i := naminginvoker.New(shared.LocalHost, shared.NamingPort)
	go i.Invoke()
}

func calcServer() {
	// Create proxy of messagingservice subscriber
	naming := namingproxy.New(shared.LocalHost, shared.NamingPort)

	// Create invoker of messagingservice
	calcInv := calculadorainvoker.New(shared.LocalHost, shared.CalculadoraPort)

	// Register Calculadora in Naming subscriber
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
