package main

import "fmt"

func main() {

	s1 := "casa"
	b := []byte(s1) // serializacao
	s2 := string(b) // desserializacao

	fmt.Println(s1, b, s2)
}
