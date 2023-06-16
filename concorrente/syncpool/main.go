package main

import (
	"fmt"
	"sync"
)

type Pessoa struct {
	Nome string
}

// Initialização do pool
var poolPessoas = sync.Pool{
	New: func() interface{} { return new(Pessoa) }}

func main() {

	poolPessoas.Put(&Pessoa{Nome: "Ana"})
	poolPessoas.Put(&Pessoa{Nome: "Bela"})
	poolPessoas.Put(&Pessoa{Nome: "Clara"})
	poolPessoas.Put(&Pessoa{Nome: "Dora"})

	p1 := poolPessoas.Get().(*Pessoa)
	p2 := poolPessoas.Get().(*Pessoa)
	p3 := poolPessoas.Get().(*Pessoa)
	p4 := poolPessoas.Get().(*Pessoa)
	p5 := poolPessoas.Get().(*Pessoa)

	fmt.Println(p1.Nome, p2.Nome, p3.Nome, p4.Nome, p5.Nome)
}
