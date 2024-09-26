package main

import (
	"fmt"
	"test/mymiddleware/distribution/invokers/calculadora/calculadorainvoker"
)

func main() {
	// start calculadora invoker
	invoker := calculadorainvoker.Invoker{}

	invoker.Invoke("localhost", 1313)
}

func process(b []byte) []byte {

	fmt.Println(b)

	r := []byte{}
	for i := 0; i < len(b); i++ {
		r = append(r, b[len(b)-i])
	}

	return r
}
