package main

import (
	"distribuida/calculadora/mymiddleware/client/proxies"
	"distribuida/calculadora/mymiddleware/server/invoker"
	"fmt"
	"mymiddleware/services/naming/proxy"
)

func main() {

	// create a built-in proxy of naming service
	namingProxy := proxy.NamingProxy{}

	// create a proxy of calculator service
	//calculator := proxies.NewCalculatorProxy()
	//converter := proxies.NewConverterProxy()

	// register service in the naming service
	namingProxy.Register("Calculator", proxies.NewCalculatorProxy())
	//namingProxy.Register("Converter", converter)

	// control loop passed to middleware
	fmt.Println("Calculator Server running!!")
	calculatorInvoker := invoker.NewCalculatorInvoker()
	//converterInvoker := invoker.NewConverter()

	go calculatorInvoker.Invoke()
	//go converterInvoker.Invoke()

	fmt.Scanln()
}

