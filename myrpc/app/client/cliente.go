package main

import (
	"fmt"
	"os"
	"test/myrpc/distribution/proxies/calculadora"
	namingproxy "test/myrpc/services/naming/proxy"
	"test/shared"
	"time"
)

func main() {
	Cliente()
}

func Cliente() {

	ClientePerf()
	os.Exit(0)

	// Obtain proxies
	naming := namingproxy.New(shared.LocalHost, shared.NamingPort)
	calc := calculadoraproxy.New(naming.Find("Calculadora"))
	fmt.Println("Cliente:: HERE", calc)

	// Chamada remota a Calculadora
	fmt.Println(calc.Som(1, 2))
}

func ClientePerf() {
	naming := namingproxy.New(shared.LocalHost, shared.NamingPort)
	calc := calculadoraproxy.New(naming.Find("Calculadora"))

	for i := 0; i < shared.StatisticSample; i++ {
		t1 := time.Now()
		for j := 0; j < shared.SampleSize; j++ {
			fmt.Println("Cliente", calc.Som(1, 2))
		}
		fmt.Println(i, ";", time.Now().Sub(t1).Milliseconds())
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("Experiemnt finalised...")
}
