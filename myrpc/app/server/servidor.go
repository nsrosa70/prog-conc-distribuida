package main

import (
	"test/myrpc/distribution/invokers/calculadora"
	namingproxy "test/myrpc/services/naming/proxy"
	"test/shared"
)

func main() {
	// Obtain proxies
	naming := namingproxy.New(shared.LocalHost, shared.NamingPort)

	// Create instance of invokers
	calcInv := calculadorainvoker.New(shared.LocalHost, shared.CalculadoraPort)

	// Register services in Naming
	naming.Bind("Calculadora", shared.NewIOR(calcInv.Ior.Host, calcInv.Ior.Port))

	// Invoke services
	calcInv.Invoke()
}
