package main

import (
	"fmt"
	"mymiddleware/services/naming/invoker"
)

func main() {

	fmt.Println("Naming servidor running!!")

	// control loop passed to invoker
	namingInvoker := invoker.NamingInvoker{}
	namingInvoker.Invoke()
}
