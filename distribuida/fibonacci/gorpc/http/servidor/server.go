package main

import (
	fibonacci "aulas/distribuida/fibonacci/impl"
	"aulas/distribuida/shared"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"strconv"
)

func servidor() {

	// cria instância da calculadora
	fibonacci := new(fibonacci.FibonacciRPC)

	// cria um novo consumer RPC e registra o fibonacci
	server := rpc.NewServer()
	err := server.RegisterName("Fibonacci", fibonacci)
	shared.ChecaErro(err, "Error ao registrar o serviço Fibonacci...")

	// // cria um listener tcp (fibonacci)
	l, err := net.Listen("tcp", ":"+strconv.Itoa(shared.FibonacciPort))
	shared.ChecaErro(err, "Servidor não inicializado")

	// associa um handler HTTP ao consumer (Fibonacci)
	server.HandleHTTP("/", "/debug")

	// aguarda por invocações
	fmt.Println("Servidor está pronto (RPC-HTTP)...")
	http.Serve(l, nil)
}

func main() {

	go servidor()

	_, _ = fmt.Scanln()
}
