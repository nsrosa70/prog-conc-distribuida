package main

import (
	"distribuida/calculadora/mymiddleware/client/proxies"
	"fmt"
	"mymiddleware/services/naming/proxy"
	"time"
)

func ExecuteExperiment() {
	// create a built-in proxy of naming service
	namingService := proxy.NamingProxy{} // cria um proxy - serviço do middleware

	// look for a service in naming service
	calculator := namingService.Lookup("Calculator").(proxies.CalculatorProxy) // cria proxy

	// invoke remote operation
	t1 := time.Now()
	for i := 0; i < 1000; i++ {
		//fmt.Print(calculator.Add(1, 2))
		calculator.Add(1, 2) // Call.service("Calculator.Add",1,2) calculator.Add(ctx,1,2)
		//fmt.Println(time.Now().Sub(t1).Milliseconds())
	}
	fmt.Println("Finished in ", time.Now().Sub(t1))
}

func main() {
	go ExecuteExperiment()

	fmt.Scanln()
}
