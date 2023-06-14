package main

import (
	"fmt"
)

func main() {
	x := 0
	go func() {
		x++
	}()
	fmt.Println(x)
}
