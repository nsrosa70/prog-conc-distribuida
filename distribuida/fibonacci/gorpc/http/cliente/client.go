package main

import (
	"aulas/distribuida/shared"
	"fmt"
	"net/rpc"
	"strconv"
)

func Cliente() {
	var reply int

	// conecta ao servidor (Fibonacci)
	clientFibo, err := rpc.DialHTTP("tcp", ":"+strconv.Itoa(shared.FIBONACCI_PORT))
	shared.ChecaErro(err, "Não foi possível criar uma conexão com o servidor Fibonacci...")

	defer func(clientFibo *rpc.Client) {
		var err = clientFibo.Close()
		shared.ChecaErro(err, "Erro ao fechar a conexão com o servidor Fibonacci...")
	}(clientFibo)

	// invoca operação remota do fibonacci
	n := 10
	err = clientFibo.Call("Fibonacci.Fibo", n, &reply)
	shared.ChecaErro(err, "Erro na invocação remota do servidor Fibonacci...")
	fmt.Printf("Fibo(%v) = %v \n", n, reply)
}

func main() {

	go Cliente()

	fmt.Scanln()
}
