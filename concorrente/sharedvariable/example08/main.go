package main

import (
	"fmt"
	"runtime"
)

func main() {

	runtime.GOMAXPROCS(2) // força o SO a executar o código Go em 'go_max_procs" threads do SO

	for i := 0; i < 10000000; i++ {
		go fmt.Print(0)
		fmt.Print(1)
	}
}
