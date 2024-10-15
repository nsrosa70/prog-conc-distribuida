package main

import (
	fibonacci "aulas/distribuida/fibonacci/impl"
	"aulas/distribuida/shared"
	"fmt"
	"net"
	"net/rpc"
	"strconv"
)

func servidor() {

	// cria instância da messagingservice
	fibonacci := new(fibonacci.FibonacciRPC)

	// cria um novo servidor RPC e registra o fibonacci
	server := rpc.NewServer()
	err := server.RegisterName("Fibonacci", fibonacci)
	shared.ChecaErro(err, "Não foi possível registrar o Fibonacci no servidor...")

	// cria um listener TCP
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(shared.FibonacciPort))
	shared.ChecaErro(err, "Não foi possível criar um listener para o Fibonacci...")

	defer func(ln net.Listener) {
		var err = ln.Close()
		shared.ChecaErro(err, "Não foi possível fechar o listener do Fibonacci...")
	}(ln)

	// aguarda por invocações
	fmt.Println("Servidor está pronto (RPC-TCP)...")
	server.Accept(ln)
}

func main() {

	go servidor()

	_, _ = fmt.Scanln()
}
