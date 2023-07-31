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
	client, err := rpc.Dial("tcp", ":"+strconv.Itoa(shared.FIBONACCI_PORT))
	shared.ChecaErro(err, "Não foi possível criar uma conexão TCP para o servidor Fibonacci...")

	defer func(client *rpc.Client) {
		var err = client.Close()
		shared.ChecaErro(err, "Não foi possível fechar a conexão TCP com o servidor Fibonacci...")
	}(client)

	// invoca operação remota do Fibonacci
	n := 10
	err = client.Call("Fibonacci.Fibo", n, &reply)
	shared.ChecaErro(err, "Erro na invocação remota do Fibonacci...")
	fmt.Printf("Fibo(%v) = %v \n", n, reply)
}

func main() {

	go Cliente()

	fmt.Scanln()
}
