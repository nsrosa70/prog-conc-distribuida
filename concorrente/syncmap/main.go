package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	var precos = new(sync.Map)
	var v interface{} = nil
	var ok = false

	precos.Store("tomate", 3)
	precos.Store("banana", 4)
	precos.Store("laranja", 5)

	key := "tomate"
	v, ok = precos.Load(key)
	if ok {
		println("Preco do",key,"=", v.(int))
	} else {
		fmt.Println("Erro: Produto não encontrado")
		os.Exit(0)
	}
	precos.Delete("laranja")
}

